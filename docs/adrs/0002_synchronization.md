# 2025-05-07 Implement Filesystem←→DB synchronization and do it weel

I haven't found a fast enough and customizable way load my notes into the database. My initial test show, that well optimized Go utility loads all of my ~4000 markdown notes in 2 seconds, which is acceptable to run even on each server start. And I can even drop all tables each time and not worry about database being out of date.

The source of truth is always the filesystem.

Markdowndb took 20s.
MDQ might be required if I would want to properly parse all of the contents.
