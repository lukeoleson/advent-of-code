#!/bin/zsh

awk ' /^$/ { print; } /./ { printf("%s ", $0); } ' input.txt |

# byr (Birth Year) - four digits; at least 1920 and at most 2002.
egrep '(.*byr:(?:19[2-9][0-9]|200[0-2])\s)' |

# iyr (Issue Year) - four digits; at least 2010 and at most 2020.
egrep '(.*iyr:(?:20(?:1[0-9]|20))\s)' |

# eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
egrep '(.*eyr:(?:20(?:2[0-9]|30))\s)' |

# hgt (Height) - a number followed by either cm or in:
# If cm, the number must be at least 150 and at most 193.
# If in, the number must be at least 59 and at most 76.
egrep '(.*hgt:(?:1(?:[5-8][0-9]|9[0-3])cm|(?:59|6[0-9]|7[0-6])in)\s)' |

# hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
egrep '(.*hcl:#[a-f0-9]{6}\s)'  |

# ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
egrep '(.*ecl:(?:(amb|blu|brn|gry|grn|hzl|oth))\s)' |

# pid (Passport ID) - a nine-digit number, including leading zeroes.
egrep -c '(.*pid:[0-9]{9}\s)'

