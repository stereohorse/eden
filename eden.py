#!/usr/bin/env python3

from tinydb import TinyDB

db = TinyDB('./db.json')
db.insert({
    'data': 'something',
    'tags': ['test', 'garbage']
})
