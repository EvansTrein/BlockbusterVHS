from config import db


class VhsTape(db.Model):
    id_num = db.Column(db.Integer, primary_key=True)
    name_film = db.Column(db.String(100), nullable=False)
    year = db.Column(db.Integer, nullable=False)
    age_rating = db.Column(db.String(3), nullable=False)
    count = db.Column(db.Integer, nullable=False)
