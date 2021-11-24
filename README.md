### database table creation


 CREATE TABLE `user_details` (
 `id` int NOT NULL AUTO_INCREMENT,
 `first_name` varchar(250) NOT NULL,
 `last_name` varchar(250) NOT NULL,
 `adress` text NOT NULL,
 `age` int NOT NULL,
 `created_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci


#### Virtualhost

<VirtualHost *:80>
ProxyPreserveHost On
ProxyRequests Off
ServerName meghacheckapi.salesjio.com
ProxyPass / http://localhost:8086/
ProxyPassReverse / http://localhost:8086/
</VirtualHost>
 

#### Extra installation
add proper configurations in src/config/config.json file
#### run and build the fiile

run 
sh buildfile.sh 

    





	
