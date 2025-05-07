# 2025-05-07 Resolve conflicts based on md5 hashes

After implementing initial sync mechanism I realized that it can stuck in the infinite loop when file is updated on the disk and then in the database and post-update hook updates makes changes to the file on the disk.

Solution – md5 hashes. I store hash of the content and check it with the contents of the changed file.

More advanced synchronization techniques might be required in the future:

- timestamps both in the database and inside file attributes
- structural diffing
- CRDTS
- completely different solution, like sqlitefs
