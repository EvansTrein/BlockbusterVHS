from dataclasses import dataclass



@dataclass
class VhsTape:
    id: int
    name_film: str
    year: int
    age_rating: str
    count: int


