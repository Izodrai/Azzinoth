-- rambler up

CREATE TABLE IF NOT EXISTS services_for_data_servers (
    `service_reference`  VARCHAR(50)     NOT NULL,
    `server_reference`   VARCHAR(50)     NOT NULL,
    
    PRIMARY KEY (service_reference, server_reference),
    CONSTRAINT services_for_data_servers_fk_service_reference FOREIGN KEY service_reference(service_reference) REFERENCES services(reference),
    CONSTRAINT services_for_data_servers_fk_server_reference FOREIGN KEY server_reference(server_reference) REFERENCES servers(reference)
) ENGINE=innodb DEFAULT CHARSET=utf8;

-- rambler down

DROP TABLE services_for_data_servers;