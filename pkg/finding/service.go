package finding

import (
	"fmt"
	"log"
	"math"
	"strings"

	model "github.com/jlopez34/meli-app/pkg/model"
)

const EPSILON float32 = 0.000001
const KENOBI_X float32 = -2.0
const KENOBI_Y float32 = 0.0
const SKYWALKER_X float32 = 1.0
const SKYWALKER_Y float32 = 0.0
const SATO_X float32 = 0.0
const SATO_Y float32 = 4.0

type Service interface {
	FindLocation(model.TransmissionRequest) (model.TransmissionResponse, error)
}

type Repository interface {
	FindLocation(model.TransmissionRequest) (model.TransmissionResponse, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) FindLocation(t model.TransmissionRequest) (model.TransmissionResponse, error) {
	var transmissionResponse model.TransmissionResponse
	var position model.Position
	var messageFinal string
	var x, y float32

	var distanceKenobi, distanceSkywalker, distanceSato float32
	var msgKenobi, msgSkywalker, msgSato []string

	for _, satellite := range t.Satellites {
		var nameSatellite = satellite.Name
		switch nameSatellite {
		case "Kenobi":
			distanceKenobi = satellite.Distance
			msgKenobi = satellite.Message
		case "Skywalker":
			distanceSkywalker = satellite.Distance
			msgSkywalker = satellite.Message
		case "Sato":
			distanceSato = satellite.Distance
			msgSato = satellite.Message
		}
	}

	if distanceKenobi == 0 || distanceSkywalker == 0 || distanceSato == 0 {
		log.Fatal("Don't possible to calculate source message")
	} else {
		x, y = GetLocation(distanceKenobi, distanceSkywalker, distanceSato)
		position.X = x
		position.Y = y
	}

	messageFinal = GetMessage(msgKenobi, msgSkywalker, msgSato)

	if len(strings.TrimSpace(messageFinal)) != 0 {
		log.Fatal("Don't possible to build message send")
	}

	transmissionResponse = model.TransmissionResponse{Position: &position, Message: messageFinal}
	return transmissionResponse, nil
}

func GetLocation(distances ...float32) (x, y float32) {
	var position model.Position

	/* dx and dy are the vertical and horizontal distances between * the circle centers. */
	var a, dx, dy, d, d1, d2, h, rx, ry float32
	var point2_x, point2_y float32

	/* Determine the straight-line distance between the centers. */
	dx = SKYWALKER_X - KENOBI_X
	dy = SKYWALKER_Y - KENOBI_Y

	/* Check for solubility. */
	d = float32(math.Sqrt(float64((dy * dy) + (dx * dx))))

	/* no solution. circles do not intersect. */
	if d > (distances[0] + distances[1]) {
		return 0, 0
	}

	/* no solution. one circle is contained in the other */
	if d < float32(math.Abs(float64(distances[0]-distances[1]))) {
		return 0, 0
	}

	/* 'point 2' is the point where the line through the circle * intersection points crosses the line between the circle * centers. */
	/* Determine the distance from point 0 to point 2. */
	a = ((distances[0] * distances[0]) - (distances[1] * distances[1]) + (d * d)) / (2.0 * d) /* Determine the coordinates of point 2. */
	point2_x = KENOBI_X + (dx * a / d)
	point2_y = KENOBI_Y + (dy * a / d)

	/* Determine the distance from point 2 to either of the * intersection points. */
	h = float32(math.Sqrt(float64((distances[0] * distances[0]) - (a * a))))

	/* Now determine the offsets of the intersection points from * point 2. */
	rx = -dy * (h / d)
	ry = dx * (h / d)

	/* Determine the absolute intersection points. */
	var intersectionPoint1_x = point2_x + rx
	var intersectionPoint2_x = point2_x - rx
	var intersectionPoint1_y = point2_y + ry
	var intersectionPoint2_y = point2_y - ry
	log.Println("INTERSECTION Circle1 AND Circle2: (" + fmt.Sprint(intersectionPoint1_x) + "," + fmt.Sprint(intersectionPoint1_y) + ")" + " AND (" + fmt.Sprint(intersectionPoint2_x) + "," + fmt.Sprint(intersectionPoint2_y) + ")")
	/* Lets determine if circle 3 intersects at either of the above intersection points. */
	dx = intersectionPoint1_x - SATO_X
	dy = intersectionPoint1_y - SATO_Y
	d1 = float32(math.Sqrt(float64((dy * dy) + (dx * dx))))
	dx = intersectionPoint2_x - SATO_X
	dy = intersectionPoint2_y - SATO_Y
	d2 = float32(math.Sqrt(float64((dy * dy) + (dx * dx))))
	if float32(math.Abs(float64(d1-distances[2]))) < EPSILON {
		log.Println("INTERSECTION Circle1 AND Circle2 AND Circle3: (" + fmt.Sprint(intersectionPoint1_x) + "," + fmt.Sprint(intersectionPoint1_y) + ")")
		position = model.Position{X: intersectionPoint1_x, Y: intersectionPoint1_y}
	} else if float32(math.Abs(float64(d2-distances[2]))) < EPSILON {
		log.Println("INTERSECTION Circle1 AND Circle2 AND Circle3: (" + fmt.Sprint(intersectionPoint2_x) + "," + fmt.Sprint(intersectionPoint2_y) + ")")
	} else {
		log.Println("INTERSECTION Circle1 AND Circle2 AND Circle3: " + " NONE")
	}

	return position.X, position.Y
}

func GetMessage(messages ...[]string) (msg string) {
	var messageFound strings.Builder
	var messageKenobi = messages[0]
	var messageSkywalker = messages[1]
	var messageSato = messages[2]

	for i := 0; i < len(messageKenobi); i++ {
		var wordKenobi = messageKenobi[i]
		if strings.EqualFold("", wordKenobi) {
			var wordSkywalker = messageSkywalker[i]
			if strings.EqualFold("", wordSkywalker) {
				messageFound.WriteString(" " + messageSato[i])
			} else {
				messageFound.WriteString(" " + wordSkywalker)
			}
		} else {
			messageFound.WriteString(" " + wordKenobi)
		}
	}

	return messageFound.String()
}
