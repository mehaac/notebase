# 2025-05-13 Synchronization – One Way Only

I realized that I have infinite loops when the file is changed in the database. So I think a logic solution is to use a one way synchronization. The ideas is that I change the files on disk only. And then files are synced to the database. I react on the filesystem events only. This way I keep the files as the source of truth and don't need hashes to check for the conflicts (see [previous ADR](./0003_synchronization_2.md)).

That in turn requires a couple of new endpoints:

- changing file's frontmatter (mostly set key and get key)
- changing files's content
