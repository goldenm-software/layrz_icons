import re


def to_kebab_case(name: str) -> str:
  return '-'.join(
    re.sub(
      r'(\s|_|-)+',
      ' ',
      re.sub(
        r'[A-Z]{2,}(?=[A-Z][a-z]+[0-9]*|\b)|[A-Z]?[a-z]+[0-9]*|[A-Z]|[0-9]+', lambda mo: ' ' + mo.group(0).lower(), name
      ),
    ).split()
  )


def to_camel_case(name: str, capitalize: bool = False) -> str:
  name = to_kebab_case(name)
  name = re.sub(r'(_|-)+', ' ', name).title().replace(' ', '')
  if capitalize:
    return name[0].upper() + name[1:]

  return name[0].lower() + name[1:]


if __name__ == '__main__':
  print(to_kebab_case('mdiIconMapping'))
  print(to_camel_case('mdiIconMapping', capitalize=True))
