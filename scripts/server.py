import argparse
import getpass
import os

class PasswordPromptAction(argparse.Action):
    def __init__(self,
             option_strings,
             dest=None,
             nargs=0,
             default=None,
             required=False,
             type=None,
             metavar=None,
             help=None):
        super(PasswordPromptAction, self).__init__(
             option_strings=option_strings,
             dest=dest,
             nargs=nargs,
             default=default,
             required=required,
             metavar=metavar,
             type=type,
             help=help)

    def __call__(self, parser, args, values, option_string=None):
        password = getpass.getpass()
        setattr(args, self.dest, password)

parser = argparse.ArgumentParser(description='Initialize server environment')
parser.add_argument('-u', dest='user',type=str, nargs=1, help='username for your MySQL server')
parser.add_argument('-p', dest='pass', action=PasswordPromptAction, type=str, required=True, help='password for your MySQL server')

args = vars(parser.parse_args())

print('Set environment variables: DB_USER, DB_PASS, DB_HOST')
os.environ['DB_USER'] = args['user'][0]
os.environ['DB_PASS'] = args['pass'][0]
os.environ['DB_HOST'] = '127.0.0.1'