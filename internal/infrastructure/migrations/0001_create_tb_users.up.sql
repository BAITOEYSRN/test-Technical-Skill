create table develop.users (
    id UUID primary key default gen_random_uuid(),
    first_name VARCHAR(100) not null,
    last_name VARCHAR(100) not null,
    date_of_birth DATE not null,
    age INT not null,
    address TEXT not null,
    created_at TIMESTAMPTZ not null default NOW(),
    updated_at TIMESTAMPTZ      
);

create index idx_users_id on develop.users (id);