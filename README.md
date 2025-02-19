# WaveKit

WaveKit is a simple on-prem tool designed to enhance observability for your RisingWave cluster, enabling faster issue detection, efficient troubleshooting, and improved performance.

WaveKit supports all RisingWave deployment types, including Docker, Kubernetes, and RisingWave Cloud.

> [!WARNING]
> **WaveKit is currently in the public preview stage.**

> [!NOTE]
> _WaveKit uses a PostgreSQL database to store key cluster metadata, including connection details like hostnames and ports for RisingWave clusters. To ensure persistence, you’ll need to self-host a PostgreSQL database to prevent metadata loss._

> [!NOTE]
> _To use WaveKit, ensure your RisingWave cluster is already running and accessible._


## Installation (Quick setup)

This method installs WaveKit with a bundled PostgreSQL database for convenience. However, if you prefer to use your own self-hosted PostgreSQL database for data persistence, skip to the next section.  

### **Starting the WaveKit Server**  

You can start the WaveKit server in two ways:  

#### **Option 1: Ephemeral Storage (No Persistence)**  
Runs WaveKit with a bundled PostgreSQL database, but metadata is stored inside the container. If the container is removed, all metadata will be lost.  

```shell
docker run --rm -p 8020:8020 --name wavekit risingwavelabs/wavekit:v0.1.2-pgbundle
```

#### **Option 2: Persistent Storage (Recommended)**  
Runs WaveKit with a bundled PostgreSQL database and stores metadata in a persistent Docker volume (`wavekit-data`), ensuring data persists across restarts.  

```shell
docker run -p 8020:8020 --name wavekit -v wavekit-data:/var/lib/postgresql risingwavelabs/wavekit:v0.1.2-pgbundle
```

### **Accessing WaveKit**  

Once the server is running, open your browser and go to:  

- **[http://localhost:8020](http://localhost:8020)**  

Use the following default credentials to log in:  
- **Username:** `root`  
- **Password:** `root`  


### Customizing WaveKit Settings

WaveKit offers flexible configuration options through either a configuration file or environment variables. For detailed information about available settings and configuration methods, please refer to our [configuration documentation](docs/config.md).

## Installation (Recommended for production)

The following section provides a step-by-step guide to setting up WaveKit with your self-hosted PostgreSQL database. This approach is recommended if you need persistent metadata with high availability.

First, create `docker-compose.yaml` file with the following content:

```yaml
version: "3.9"
services:
  wavekit:
    image: cloudcarver/wavekit:v0.1.2
    ports:
      - "8020:8020"
    environment:
      WK_PORT: 8020
      WK_PG_HOST: localhost
      WK_PG_PORT: 5432
      WK_PG_USER: postgres
      WK_PG_PASSWORD: postgres
      WK_PG_DB: postgres
      WK_JWT_SECRET: 9138e41195112b568e22480f18a42dd69b38fab5ee1a36fbf63d49b22097d22a
      WK_ROOT_PASSWORD: '123456'
      WK_RISECTLDIR: /

  db: 
    image: "postgres:latest"
    ports:
      - "5432:5432"
    environment:
      POSTGRES_PASSWORD: postgres
    volumes:
      - db-data:/var/lib/postgresql/data

  rw:
    image: risingwavelabs/risingwave:v2.1.2
    ports:
      - 4566:4566
      - 5690:5690
      - 5691:5691
    command: "standalone --meta-opts=\" \
                    --listen-addr 0.0.0.0:5690 \
                    --advertise-addr rw:5690 \
                    --dashboard-host 0.0.0.0:5691 \
                    --prometheus-host 0.0.0.0:1250 \
                    --backend sqlite  \
                    --sql-endpoint /root/single_node.db \
                    --state-store hummock+fs:///root/state_store \
                    --data-directory hummock_001\" \
                 --compute-opts=\" \
                    --listen-addr 0.0.0.0:5688 \
                    --prometheus-listener-addr 0.0.0.0:1250 \
                    --advertise-addr rw:5688 \
                    --async-stack-trace verbose \
                    --parallelism 4 \
                    --total-memory-bytes 2147483648 \
                    --role both \
                    --meta-address http://0.0.0.0:5690\" \
                 --frontend-opts=\" \
                   --listen-addr 0.0.0.0:4566 \
                   --advertise-addr rw:4566 \
                   --prometheus-listener-addr 0.0.0.0:1250 \
                   --health-check-listener-addr 0.0.0.0:6786 \
                   --meta-addr http://0.0.0.0:5690 \
                   --frontend-total-memory-bytes=500000000\" \
                 --compactor-opts=\" \
                   --listen-addr 0.0.0.0:6660 \
                   --prometheus-listener-addr 0.0.0.0:1250 \
                   --advertise-addr rw:6660 \
                   --meta-address http://0.0.0.0:5690 \
                   --compactor-total-memory-bytes=1000000000\""

volumes:
  db-data:
  wavekit-data:

```

Start WaveKit by running the following command:

```shell
docker compose up
```

When the server is running, open your browser and go to: **[http://localhost:8020](http://localhost:8020)**  


## Contributing to WaveKit

We welcome contributions to WaveKit! Please refer to our [CONTRIBUTING.md](CONTRIBUTING.md) for more information on how to contribute to the project.


## WaveKit Editions

WaveKit is available in two editions:  

- **WaveKit-Lite** – A lightweight, open-source edition that includes core functionalities. Licensed under Apache 2.0.  
- **WaveKit-Pro** – A full-featured edition with advanced capabilities. A license key is required for access. To apply, contact us at [sales@risingwave-labs.com](mailto:sales@risingwave-labs.com) or [fill out this form](https://cloud.risingwave.com/auth/license_key/).
