package main

func maxNumberOfFamilies(n int, reservedSeats [6][2]int) int {
    res := 0
    datas := make([][]bool,0)
    for i := 0;i <n;i++ {
        datas = append(datas,make([]bool,10))
    }
    
    for i := 0 ; i < len(reservedSeats) ; i++ {
        x := reservedSeats[i][0] -1 
        y := reservedSeats[i][1] -1
        
        datas[x][y] = true
    }
    
    for i := 0;i < len(datas);i++ {
        count := 0
        for j := 0; j < len(datas[0]);j++ {
            if datas[i][j] == true {
                count = 0
            } else {
                count ++
            }
            
            if count == 4 {
                res ++
                count = 0
            }
            
            if j == 2 && count == 3 {
                count = 2
			}
			
			if j == 2 && count == 1 {
				count = 0
			}
            
            if j == 8{
                break
            }
        }
    }
    
    return res
}

func main() {
//3, reservedSeats = [[1,2],[1,3],[1,8],[2,6],[3,1],[3,10]]
	b := [6][2]int{{1, 2}, {1,3}, {1,8},{2,6},{3,1},{3,10}}

	println(maxNumberOfFamilies(3,b))
}