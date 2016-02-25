psql postgres -c "CREATE USER mybb_gateway_user WITH PASSWORD 'mybb_gateway_password' SUPERUSER CREATEDB CREATEROLE INHERIT LOGIN"
psql postgres -c "CREATE DATABASE mybb_gateway_db OWNER mybb_gateway_user ENCODING 'UTF8' TEMPLATE template0"
