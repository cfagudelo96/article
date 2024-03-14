create table if not exists restaurant(
    id uuid primary key,
    name text not null,
    schedule jsonb not null,
    created_at timestamptz default now() not null,
    updated_at timestamptz default now() not null,
    deactivaded boolean default false not null
);
