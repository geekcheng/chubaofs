Object Subsystem (ObjectNode)
==============================

How To Provide Object Storage Service with Object Subsystem (ObjectNode)
-------------------------------------------------------------------------

Start a ObjectNode process by execute the server binary of ChubaoFS you built with ``-c`` argument and specify configuration file.

.. code-block:: bash

   nohup cfs-server -c objectnode.json &


Configurations
-----------------------
Object Node using `JSON` format configuration file.


**Properties**

.. csv-table::
   :header: "Key", "Type", "Description", "Mandatory"

   "role", "string", "Role of process and must be set to *objectnode*", "Yes"
   "listen", "string", "Listen and accept port of the server. Default: 80", "Yes"
   "region", "string", "Region of this gateway. Used by S3-like interface signature validation. Default: cfs_default", "No"
   "domains", "string slice", "
   | Format: *DOMAIN*.
   | DOMAIN: Domain of S3-like interface which makes wildcard domain support", "No"
   "logDir", "string", "Log directory", "Yes"
   "logLevel", "string", "Level operation for logging. Default is *error*", "No"
   "masters", "string slice", "
   | Format: *HOST:PORT*.
   | HOST: Hostname, domain or IP address of master (resource manager).
   | PORT: port number which listened by this master", "Yes"
   "exporterPort", "string", "Port for monitor system", "No"
   "prof", "string", "Pprof port", "Yes"


**Example:**

.. code-block:: json

   {
        "role": "objectnode",
        "listen": 80,
        "region": "test",
        "domains": [
            "object.cfs.local"
        ],
        "logDir": "/opt/cfs/objectnode/logs",
        "logLevel": "debug",
        "masters": [
	        "172.20.240.95:7002",
	        "172.20.240.94:7002",
	        "172.20.240.67:7002"
        ],
        "exporterPort": 9512,
        "prof": "7013"
   }

Fetch Authentication Keys
----------------------------
Authentication keys owned by volume and stored with volume view (volume topology) by Resource Manager (Master).
User can fetch it by using administration API, see **Get Volume Information** at :doc:`/admin-api/master/volume`

Using Object Storage Interface
-------------------------------
Object Subsystem (ObjectNode) provides S3-compatible object storage interface, so that you can operate files by using native Amazon S3 SDKs.

For detail about list of supported APIs, see **Supported S3-compatible APIs** at :doc:`/design/objectnode`

For detail about list of supported SDKs, see **Supported SDKs** at :doc:`/design/objectnode`
