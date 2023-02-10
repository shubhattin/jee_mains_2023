from pydantic.dataclasses import dataclass

@dataclass
class Jee():
    not_found: str

@dataclass
class LangDBModel():
    jee: Jee
