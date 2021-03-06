// Copyright 2018 The ChubaoFS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package objectnode

import (
	"encoding/xml"
	"net/http"
	"strings"

	"github.com/chubaofs/chubaofs/util/log"
)

type ErrorCode struct {
	ErrorCode    string
	ErrorMessage string
	StatusCode   int
}

func (code ErrorCode) ServeResponse(w http.ResponseWriter, r *http.Request) error {
	var err error
	var marshaled []byte
	var xmlError = struct {
		XMLName   xml.Name `xml:"Error"`
		Code      string   `xml:"Code"`
		Message   string   `xml:"Message"`
		Resource  string   `xml:"Resource"`
		RequestId string   `xml:"RequestId"`
	}{
		Code:      code.ErrorCode,
		Message:   code.ErrorMessage,
		Resource:  r.URL.String(),
		RequestId: RequestIDFromRequest(r),
	}
	if marshaled, err = xml.Marshal(&xmlError); err != nil {
		return err
	}
	w.Header().Set(HeaderNameContentType, HeaderValueContentTypeXML)
	w.WriteHeader(code.StatusCode)
	log.LogInfof("Error info : %s", string(marshaled))
	if _, err = w.Write(marshaled); err != nil {
		return err
	}
	return nil
}

func ServeInternalStaticErrorResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(HeaderNameContentType, HeaderValueContentTypeXML)
	w.WriteHeader(http.StatusInternalServerError)
	sb := strings.Builder{}
	sb.WriteString(xml.Header)
	sb.WriteString("<Error><Code>InternalError</Code><Message>We encountered an internal error. Please try again.</Message><Resource>")
	sb.WriteString(r.URL.String())
	sb.WriteString("</Resource><RequestId>")
	sb.WriteString(RequestIDFromRequest(r))
	sb.WriteString("</RequestId></Error>")
	_, _ = w.Write([]byte(sb.String()))
}

var (
	UnsupportedOperation                = ErrorCode{ErrorCode: "UnsupportedOperation", ErrorMessage: "Operation is not supported", StatusCode: http.StatusBadRequest}
	AccessDenied                        = ErrorCode{ErrorCode: "AccessDenied", ErrorMessage: "Access Denied", StatusCode: http.StatusForbidden}
	BadDigest                           = ErrorCode{ErrorCode: "BadDigest", ErrorMessage: "The Content-MD5 you specified did not match what we received.", StatusCode: http.StatusBadRequest}
	BucketNotExisted                    = ErrorCode{ErrorCode: "BucketNotExisted", ErrorMessage: "The requested bucket name is not existed.", StatusCode: http.StatusNotFound}
	BucketNotExistedForHead             = ErrorCode{ErrorCode: "BucketNotExisted", ErrorMessage: "The requested bucket name is not existed.", StatusCode: http.StatusConflict}
	BucketNotEmpty                      = ErrorCode{ErrorCode: "BucketNotEmpty", ErrorMessage: "The bucket you tried to delete is not empty.", StatusCode: http.StatusConflict}
	BucketNotOwnedByYou                 = ErrorCode{ErrorCode: "BucketNotOwnedByYou", ErrorMessage: "The bucket is not owned by you.", StatusCode: http.StatusConflict}
	KeyTooLongError                     = ErrorCode{ErrorCode: "KeyTooLongError", ErrorMessage: "", StatusCode: http.StatusBadRequest}
	InvalidKey                          = ErrorCode{ErrorCode: "InvalidKey", ErrorMessage: "Object key is Illegal", StatusCode: http.StatusBadRequest}
	EntityTooSmall                      = ErrorCode{ErrorCode: "EntityTooSmall", ErrorMessage: "Your proposed upload is smaller than the minimum allowed object size.", StatusCode: http.StatusBadRequest}
	EntityTooLarge                      = ErrorCode{ErrorCode: "EntityTooLarge", ErrorMessage: "Your proposed upload exceeds the maximum allowed object size.", StatusCode: http.StatusBadRequest}
	IncorrectNumberOfFilesInPostRequest = ErrorCode{ErrorCode: "IncorrectNumberOfFilesInPostRequest", ErrorMessage: "POST requires exactly one file upload per request.", StatusCode: http.StatusBadRequest}
	InternalError                       = ErrorCode{ErrorCode: "InternalError", ErrorMessage: "We encountered an internal error. Please try again.", StatusCode: http.StatusInternalServerError}
	InvalidArgument                     = ErrorCode{ErrorCode: "InvalidArgument", ErrorMessage: "Invalid Argument", StatusCode: http.StatusBadRequest}
	InvalidBucketName                   = ErrorCode{ErrorCode: "InvalidBucketName", ErrorMessage: "The specified bucket is not valid.", StatusCode: http.StatusBadRequest}
	InvalidRange                        = ErrorCode{ErrorCode: "InvalidRange", ErrorMessage: "The requested range cannot be satisfied.", StatusCode: http.StatusRequestedRangeNotSatisfiable}
	MissingContentLength                = ErrorCode{ErrorCode: "MissingContentLength", ErrorMessage: "You must provide the Content-Length HTTP header.", StatusCode: http.StatusLengthRequired}
	NoSuchBucket                        = ErrorCode{ErrorCode: "NoSuchBucket", ErrorMessage: "The specified bucket does not exist.", StatusCode: http.StatusNotFound}
	NoSuchKey                           = ErrorCode{ErrorCode: "NoLoggingStatusForKey", ErrorMessage: "The specified key does not exist.", StatusCode: http.StatusNotFound}
	PreconditionFailed                  = ErrorCode{ErrorCode: "PreconditionFailed", ErrorMessage: "At least one of the preconditions you specified did not hold.", StatusCode: http.StatusPreconditionFailed}
	MaxContentLength                    = ErrorCode{ErrorCode: "MaxContentLength", ErrorMessage: "Content-Length is bigger than 20KB.", StatusCode: http.StatusLengthRequired}
)
