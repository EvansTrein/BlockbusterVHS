from config import db


class VhsTape(db.Model):
    __bind_key__ = 'rentals'
    __tablename__ = 'vhs_tapes'
    id_num = db.Column(db.Integer, primary_key=True)
    title = db.Column(db.String)
    year = db.Column(db.Integer)
    age_rating = db.Column(db.String)
    count = db.Column(db.Integer)


class Client(db.Model):
    __bind_key__ = 'rentals'
    __tablename__ = 'clients'
    id_client = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String)
    age = db.Column(db.Integer)
    phone = db.Column(db.String)

    def __repr__(self) -> str:
        return f'{self.name}'


class Rental(db.Model):
    __bind_key__ = 'rentals'
    __tablename__ = 'rentals'
    id_link = db.Column(db.Integer, primary_key=True)
    client_id = db.Column(db.Integer, db.ForeignKey('clients.id_client'))
    vhs_tape_id = db.Column(db.Integer, db.ForeignKey('vhs_tapes.id_num'))
    client = db.relationship('Client', backref=db.backref('rentals', lazy='dynamic'))
    vhs_tapes = db.relationship('VhsTape', backref=db.backref('rentals', lazy='dynamic'))