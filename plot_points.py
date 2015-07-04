#!/usr/local/bin/python
import csv
import click
import numpy as np
import matplotlib.pyplot as plt
from dateutil import parser

@click.command()
@click.argument('filename')
@click.option('--axis', default=1, help="1 = x-axis, 2 = y-axis")
def plot(filename, axis):
    from_csv = list(csv.reader(open(filename)))
    data = np.asarray(sanitize_data(from_csv))

    plt.plot(data[..., 0], data[..., int(axis)])
    plt.show()

def sanitize_data(data):
    def _parse_time(timestr):
        timezone_removed = ' '.join(timestr.split()[:2])
        return parser.parse(timezone_removed).strftime('%M%S%f')
    return [(_parse_time(x), float(y), float(z)) for x, y, z in data]

if __name__ == '__main__':
    plot()
