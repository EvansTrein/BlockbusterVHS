from flask import Flask
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)
app.secret_key = 'my_secret_key_1234567890'

app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///database.db"
app.config['SQLALCHEMY_BINDS'] = {'clients': 'sqlite:///clients.db',
                                'rentals': 'sqlite:///rentals.db'}

db = SQLAlchemy(app)