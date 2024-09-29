from config import db


class VhsTape(db.Model):
    id_num = db.Column(db.Integer, primary_key=True)
    title = db.Column(db.String)
    year = db.Column(db.Integer)
    age_rating = db.Column(db.String)
    count = db.Column(db.Integer)


class Client(db.Model):
    __bind_key__ = 'clients'
    id_client = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String)
    age = db.Column(db.Integer)
    phone = db.Column(db.Integer)

