package piscine

func PodiumPosition(podium [][]string) [][]string {
	// for i := 0; i < len(podium)-1; i++ {
	// 	for j := 0; j < len(podium)-i-1; j++ {
	// 		if podium[j][0] > podium[j+1][0] {
	// 			podium[j][0], podium[j+1][0] = podium[j+1][0], podium[j][0]
	// 		}
	// 	}
	// }
	for i := 0; i < len(podium)/2; i++ {
		podium[i][0], podium[len(podium)-1-i][0] = podium[len(podium)-1-i][0], podium[i][0]
	}
	return podium
}
