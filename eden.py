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
    ret_container = []
    try:
        curses.wrapper(get_with_ui, ret_container)
    except KeyboardInterrupt:
        pass

    if len(ret_container):
        print(ret_container[0])


def get_with_ui(stdscr, ret_container):
    curses.curs_set(0)
    curses.nonl()

    inp_win = curses.newwin(1, curses.COLS - 1,
                            curses.LINES - 1, 0)

    results_win = curses.newwin(curses.LINES - 2,
                                curses.COLS - 1,
                                0, 0)

    def dsr(rets, sel_i):
        draw_search_rets(results_win, rets, sel_i)

    search_str = ''
    search_rets = []
    selected_i = None

    ix = get_index()
    with ix.searcher() as searcher:
        while True:
            k = stdscr.getch()

            if k == ca.CR:
                if selected_i is not None:
                    ret_container.append(search_rets[selected_i]['text'])
                return

            if k == ca.LF:
                if search_rets and selected_i is not None:
                    if selected_i < (len(search_rets) - 1):
                        selected_i += 1
                    else:
                        selected_i = 0

                    dsr(search_rets, selected_i)
                continue

            if k == ca.VT:
                if search_rets and selected_i is not None:
                    if selected_i > 0:
                        selected_i -= 1
                    else:
                        selected_i = len(search_rets) - 1

                    dsr(search_rets, selected_i)
                continue

            if k == ca.DEL:
                search_str = search_str[:-1]
            elif not 32 <= k <= 126:
                continue
            else:
                search_str += chr(k)

            inp_win.erase()
            inp_win.addstr(search_str)
            inp_win.refresh()

            qp = wq.QueryParser('text', ix.schema)
            query = qp.parse(search_str)

            search_rets = []
            for ret in searcher.search(query):
                search_ret = {
                    'pos': ret.pos,
                    'text': ret.fields()['text']
                }

                search_rets.append(search_ret)

            if search_rets:
                selected_i = 0
            else:
                selected_i = None

            dsr(search_rets, selected_i)


def draw_search_rets(window, search_rets, selected_i):
    window.erase()

    for ret in search_rets:
        if selected_i == ret['pos']:
            window.addstr(ret['pos'], 0, ret['text'], curses.A_REVERSE)
        else:
            window.addstr(ret['pos'], 0, ret['text'])

    window.refresh()


if __name__ == '__main__':
    args = parse_args()
    args.handle(args)
