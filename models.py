from config import db


class VhsTape(db.Model):
    id_num = db.Column(db.Integer, primary_key=True)
    name_film = db.Column(db.String)
    year = db.Column(db.Integer)
    age_rating = db.Column(db.String)
    count = db.Column(db.Integer)
