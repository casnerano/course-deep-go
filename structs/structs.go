package structs

const (
	maxManaValue   = 1_000
	maxHealthValue = 1_000
	maxLevelValue  = 10
)

type Option func(*GamePerson)

func WithName(name string) Option {
	return func(person *GamePerson) {
		copy(person.name[:], name)
	}
}

func WithCoordinates(x, y, z int) Option {
	return func(person *GamePerson) {
		person.coords[0] = int32(x)
		person.coords[1] = int32(y)
		person.coords[2] = int32(z)
	}
}

func WithGold(gold uint32) Option {
	return func(person *GamePerson) {
		person.gold = gold
	}
}

func WithMana(mana uint16) Option {
	return func(person *GamePerson) {
		if mana <= maxManaValue {
			person.packedHealthAndMana[1] |= byte((mana >> 4) & 0x3F)
			person.packedHealthAndMana[2] = byte((mana & 0x0F) << 4)
		}
	}
}

func WithHealth(health uint16) Option {
	return func(person *GamePerson) {
		if health <= maxHealthValue {
			person.packedHealthAndMana[0] = byte(health >> 2)
			person.packedHealthAndMana[1] |= byte(health&0x03) << 6
		}
	}
}

func WithRespect(respect uint8) Option {
	return func(person *GamePerson) {
		person.packedDataAndAttributes[2] |= byte(respect&0x0F) << 3
	}
}

func WithStrength(strength uint8) Option {
	return func(person *GamePerson) {
		person.packedDataAndAttributes[1] |= byte(strength & 0x07)
		person.packedDataAndAttributes[2] |= byte((strength&0x08)>>3) << 7
	}
}

func WithExperience(experience uint8) Option {
	return func(person *GamePerson) {
		person.packedDataAndAttributes[1] |= byte(experience&0x0F) << 3
	}
}
func WithLevel(level uint8) Option {
	return func(person *GamePerson) {
		if level <= maxLevelValue {
			person.packedDataAndAttributes[0] |= byte(level&0x07) << 0
			person.packedDataAndAttributes[1] |= byte((level&0x08)>>3) << 7
		}
	}
}

func WithHouse() Option {
	return func(person *GamePerson) {
		person.packedDataAndAttributes[0] |= 1 << 5
	}
}

func WithGun() Option {
	return func(person *GamePerson) {
		person.packedDataAndAttributes[0] |= 1 << 4
	}
}

func WithFamily() Option {
	return func(person *GamePerson) {
		person.packedDataAndAttributes[0] |= 1 << 3
	}
}
func WithType(personType int) Option {
	return func(person *GamePerson) {
		var t byte
		switch personType {
		case BlacksmithGamePersonType:
			t = 1
		case WarriorGamePersonType:
			t = 2
		default:
			t = 0
		}
		person.packedDataAndAttributes[0] |= t << 6
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	gold   uint32
	coords [3]int32
	name   [42]byte

	// Packed health and mana value by bits (big-endian):
	// 00000000 00000000 00000000
	/// 1-10: health, 11-20: mana
	packedHealthAndMana [3]byte

	// Packed data and attributes by bits (big-endian):
	// 00000000 00000000 00000000
	// 1-2: type, 3: has house, 4: has gun, 5: has family
	// 6-9: level, 10-13: experience, 14-17: strength, 18-21: respect
	packedDataAndAttributes [3]byte
}

func NewGamePerson(options ...Option) GamePerson {
	gp := GamePerson{}
	for _, option := range options {
		option(&gp)
	}
	return gp
}

func (p *GamePerson) Name() string {
	return string(p.name[:])
}

func (p *GamePerson) X() int {
	return int(p.coords[0])
}

func (p *GamePerson) Y() int {
	return int(p.coords[1])
}

func (p *GamePerson) Z() int {
	return int(p.coords[2])
}

func (p *GamePerson) Gold() uint32 {
	return p.gold
}

func (p *GamePerson) Mana() uint16 {
	return (uint16(p.packedHealthAndMana[1])&0x3F)<<4 | (uint16(p.packedHealthAndMana[2]) >> 4)
}

func (p *GamePerson) Health() uint16 {
	return uint16(p.packedHealthAndMana[0])<<2 | uint16(p.packedHealthAndMana[1]>>6)
}

func (p *GamePerson) Respect() uint8 {
	return uint8(p.packedDataAndAttributes[2]>>3) & 0x0F
}

func (p *GamePerson) Strength() uint8 {
	return (p.packedDataAndAttributes[1] & 0x07) | ((p.packedDataAndAttributes[2]>>7)&0x01)<<3
}

func (p *GamePerson) Experience() uint8 {
	return uint8(p.packedDataAndAttributes[1]>>3) & 0x0F
}

func (p *GamePerson) Level() uint8 {
	return (p.packedDataAndAttributes[0] & 0x07) | ((p.packedDataAndAttributes[1]>>7)&0x01)<<3
}

func (p *GamePerson) HasHouse() bool {
	return p.packedDataAndAttributes[0]&(1<<5) != 0
}

func (p *GamePerson) HasGun() bool {
	return p.packedDataAndAttributes[0]&(1<<4) != 0
}

func (p *GamePerson) HasFamily() bool {
	return p.packedDataAndAttributes[0]&(1<<3) != 0
}

func (p *GamePerson) Type() int {
	switch p.packedDataAndAttributes[0] >> 6 {
	case 1:
		return BlacksmithGamePersonType
	case 2:
		return WarriorGamePersonType
	default:
		return BuilderGamePersonType
	}
}
