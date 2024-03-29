package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = false

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var deutsch = `Vor dem Gesetz steht ein Türhüter. 
  Zu diesem Türhüter kommt ein Mann vom Lande und bittet um Eintritt in das Gesetz. 
  Aber der Türhüter sagt, daß er ihm jetzt den Eintritt nicht gewähren könne. 
  Der Mann überlegt und fragt dann, ob er also später werde eintreten dürfen.

  »Es ist möglich«, sagt der Türhüter, »jetzt aber nicht.«

  Da das Tor zum Gesetz offensteht wie immer und der Türhüter beiseite tritt, 
  bückt sich der Mann, um durch das Tor in das Innere zu sehn. Als der Türhüter das merkt, 
  lacht er und sagt:

  »Wenn es dich so lockt, versuche es doch, trotz meines Verbotes hineinzugehn. 
  Merke aber: Ich bin mächtig. Und ich bin nur der unterste Türhüter. 
  Von Saal zu Saal stehn aber Türhüter, einer mächtiger als der andere. 
  Schon den Anblick des dritten kam nicht einmal ich mehr ertragen.«`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})

	t.Run("negative test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"аbc", // 15
			}
			require.NotEqual(t, expected, Top10(text))
		}
	})

	t.Run("Deutsch test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"der",      // 8
				"türhüter", // 8
				"das",      // 5
				"und",      // 5
				"aber",     // 4
				"er",       // 3
				"es",       // 3
				"gesetz",   // 3
				"ich",      // 3
				"mann",     // 3
			}
			require.Equal(t, expected, Top10(deutsch))
		} else {
			expected := []string{
				"der",      // 8
				"das",      // 5
				"Türhüter", // 8
				"und",      // 5
				"er",       // 3
				"Eintritt", // 4
				"Gesetz",   // 3
				"Mann",     // 3
				"Saal",     // 3
				"Tor",      // 3
			}
			require.Equal(t, expected, Top10(deutsch))
		}
	})
}
