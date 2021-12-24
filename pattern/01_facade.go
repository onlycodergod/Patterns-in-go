package pattern

import (
	"fmt"
	"strings"
)

/*
	Реализовать паттерн "фасад".
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного паттерна на практике
	https://en.wikipedia.org/wiki/Facade_pattern


Реализация паттерна в коде ниже на примере музыкальной группы.
*/

// Подсистемы: вокалист и музыканты, которые игруют в группе
// Сложная подсистема состоит из множества разнообразных классов.
// Для того, чтобы заставить их что-то делать, нужно знать подробности устройства подсистемы, порядок инициализации объектов и так далее.
// Классы подсистемы не знают о существовании фасада и работают друг с другом напрямую.

type Vocalist struct {
}

type Guitarist struct {
}

type Drummer struct {
}

// Функции в подсистемах

func (v *Vocalist) Sing() string {
	return "Музыкант поёт"
}

func (g *Guitarist) PlayMelody() string {
	return "Гитарист играет мелодию"
}

func (d *Drummer) PlayBit() string {
	return "Барабанщик играет бит"
}

// Фасад: группа, в которой участвуют музыканты
// Фасад предоставляет быстрый доступ к определённой функциональности подсистемы.
// Он знает, каким классам нужно переадресовать запрос, и какие данные для этого нужны.

type MusicianGroup struct {
	vocal  *Vocalist
	guitar *Guitarist
	drum   *Drummer
}

func NewMusicianGroup() *MusicianGroup {
	return &MusicianGroup{
		vocal:  &Vocalist{},
		guitar: &Guitarist{},
		drum:   &Drummer{},
	}
}

// Реализация паттерна: вызов методов различных структур в определенном порядке
func (m *MusicianGroup) PlaySong() {
	result := []string{
		m.guitar.PlayMelody(),
		m.drum.PlayBit(),
		m.vocal.Sing(),
	}
	playStr := strings.Join(result, "\n")
	fmt.Println(playStr)
}
