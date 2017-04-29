#!/usr/bin/env python3

import argparse

import curses
import curses.ascii as ca

import os.path as op
import os

import whoosh.index as wi
import whoosh.fields as wf
import whoosh.qparser as wq


def parse_args():
    parser = argparse.ArgumentParser(prog='eden')
    subparsers = parser.add_subparsers()

    put_parser = subparsers.add_parser('put')
    put_parser.add_argument('text', nargs='+')
    put_parser.set_defaults(handle=put_text)

    get_parser = subparsers.add_parser('get')
    get_parser.set_defaults(handle=get_text)

    return parser.parse_args()


def put_text(args):
    iw = get_index().writer()
    iw.add_document(text=' '.join(args.text))
    iw.commit()


def get_index():
    home_dir = op.expanduser('~')
    storage_dir = op.join(home_dir, '.eden')

    if op.isdir(storage_dir):
        return wi.open_dir(storage_dir)

    if op.exists(storage_dir):
        raise Exception(storage_dir + ' exists')

    os.mkdir(storage_dir)

    schema = wf.Schema(text=wf.TEXT(stored=True))
    return wi.create_in(storage_dir, schema)


def get_text(args):
    try:
        curses.wrapper(init_ui)
    except KeyboardInterrupt:
        pass


def init_ui(stdscr):
    curses.curs_set(0)

    inp_win = curses.newwin(1, curses.COLS - 1,
                            curses.LINES - 1, 0)

    results_win = curses.newwin(curses.LINES - 2,
                                curses.COLS - 1,
                                0, 0)

    search_str = ''

    ix = get_index()
    with ix.searcher() as searcher:
        while True:
            k = stdscr.getch()

            if k == curses.KEY_ENTER or k == 10 or k == 13:
                return

            if k == curses.KEY_BACKSPACE or k == ca.DEL:
                search_str = search_str[:-1]
            else:
                search_str += chr(k)

            inp_win.erase()
            inp_win.addstr(search_str)
            inp_win.refresh()

            qp = wq.QueryParser('text', ix.schema)
            query = qp.parse(search_str)

            results = searcher.search(query)

            results_win.erase()
            for ret in results:
                text = ret.fields()['text']
                results_win.addstr(ret.pos, 0, text)
            results_win.refresh()


if __name__ == '__main__':
    args = parse_args()
    args.handle(args)
