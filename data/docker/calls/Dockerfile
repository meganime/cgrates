FROM debian:bullseye-slim

RUN groupadd -r mongodb && useradd -r -g mongodb mongodb
RUN groupadd -r redis && useradd -r -g redis redis 
RUN groupadd -r nats && useradd -r -g nats nats

RUN mkdir -p /usr/share/cgrates/conf /usr/share/man/man1 /var/spool/cgrates/ers/in /var/spool/cgrates/ers/out /var/spool/cgrates/cdre/csv /var/spool/cgrates/cdre/fwv /var/spool/cgrates/tpe /var/spool/cgrates/failed_posts /var/spool/cgrates/analyzers /run /data/configdb /data/db /kafka /logs /opt/nats

RUN echo 'debconf debconf/frontend select Noninteractive' | debconf-set-selections

RUN apt-get update

RUN apt-get install -y apt-utils wget gnupg gnupg2 apt-transport-https curl redis-server git gcc build-essential rsyslog procps gosu "mariadb-server" mariadb-backup socat default-jdk neovim net-tools

RUN curl -fsSL https://pgp.mongodb.com/server-7.0.asc | gpg -o /usr/share/keyrings/mongodb-server-7.0.gpg  --dearmor
RUN echo "deb [ signed-by=/usr/share/keyrings/mongodb-server-7.0.gpg ] https://repo.mongodb.org/apt/debian bullseye/mongodb-org/7.0 main" | tee /etc/apt/sources.list.d/mongodb-org-7.0.list

RUN apt-get update
RUN apt-get install -y mongodb-org postgresql --fix-missing

RUN set -ex; \
	rm -rf /var/lib/apt/lists/*; \
	rm -rf /var/lib/mysql; \
	mkdir -p /var/lib/mysql /var/run/mysqld; \
	chown -R mysql:mysql /var/lib/mysql /var/run/mysqld; \
	chmod 777 /var/run/mysqld; \
	find /etc/mysql/ -name '*.cnf' -print0 \
		| xargs -0 grep -lZE '^(bind-address|log|user\s)' \
		| xargs -rt -0 sed -Ei 's/^(bind-address|log|user\s)/#&/'; \
	echo '[mysqld]\nskip-host-cache\nskip-name-resolve' > /etc/mysql/conf.d/docker.cnf

RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN touch /logs/mariadb.log /logs/mariadb_script.log /logs/rabbitmq.log 
RUN chmod 777  /logs/mariadb.log /logs/mariadb_script.log /logs/rabbitmq.log

RUN wget -O go.tgz "https://storage.googleapis.com/golang/go1.21.0.linux-amd64.tar.gz" --progress=dot:giga
RUN tar -C /usr/local -xzf go.tgz
RUN rm go.tgz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
RUN	go version

COPY ./scripts /scripts
COPY ./scripts/service /usr/local/bin/service

COPY ./conf/rsyslogd.conf /etc/rsyslogd.conf
COPY ./conf/rsyslog.d /etc/rsyslog.d
COPY ./conf/redis.conf /etc/redis/redis.conf

COPY ./docker-entrypoint.sh /usr/local/bin/

# Set environment variables
ENV KAMAILIO_VERSION=56
ENV PJSUA_VERSION=2.9
ENV PJSUA_URL=https://github.com/pjsip/pjproject/archive/refs/tags/${PJSUA_VERSION}.tar.gz
ENV TMP_DIR=/tmp

# Install necessary dependencies
RUN apt-get update && apt-get install -y \
    libasound2-dev \
    libssl-dev 

# Add Kamailio repository key
RUN apt-key adv --fetch-keys http://deb.kamailio.org/kamailiodebkey.gpg

# Add Kamailio repository
RUN echo "deb http://deb.kamailio.org/kamailio${KAMAILIO_VERSION} bullseye main" >> /etc/apt/sources.list.d/kamailio.list

# Add Kamailio repository src
RUN echo "deb-src http://deb.kamailio.org/kamailio${KAMAILIO_VERSION} bullseye main" >> /etc/apt/sources.list.d/kamailio.list

# Install Kamailio packages
RUN apt-get update && apt-get install -y \
    kamailio \
    kamailio-extra-modules \
    kamailio-json-modules

# Clean up apt cache
RUN apt-get clean && rm -rf /var/lib/apt/lists/*

# Download and extract PJSUA
RUN mkdir -p ${TMP_DIR} && \
    curl -L ${PJSUA_URL} | tar xz -C ${TMP_DIR} --strip-components=1 && \
    cd ${TMP_DIR} && \
    ./configure && make && make install && \
    rm -rf ${TMP_DIR}

ENTRYPOINT ["docker-entrypoint.sh"]