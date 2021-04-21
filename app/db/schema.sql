-- All tx hashes generated on root chain, to be persisted in this table
create table root_chain (
    txhash char(66) primary key,
    code smallint not null,
    msg varchar not null
);

-- All tx hashes generated on child chain, to be persisted in this table
create table child_chain (
    txhash char(66) primary key,
    code smallint not null,
    msg varchar not null
);
