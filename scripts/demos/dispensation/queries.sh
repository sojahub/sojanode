#!/usr/bin/env bash
rm -rf all.json pending.json completed.json
sojanoded q dispensation records-by-name ar1 All>> all.json
sojanoded q dispensation records-by-name ar1 Pending >> pending.json
sojanoded q dispensation records-by-name ar1 Completed>> completed.json