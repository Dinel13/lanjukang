alter table users
add name_service varchar,
  add bank varchar,
  add rekening varchar;
  
create unique index users_name_service_uindex on users (name_service);