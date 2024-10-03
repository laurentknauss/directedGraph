// For simplicity, we use slices to implement the adjacent list. 


package main

import (
	"fmt"
)



// Implementing the graph structure ( graph represents an adjacency list graph) 
type Graph struct { 
			vertices []	*Vertex // a list that contains vertices . 

}



// Implementing the vertex structure ( vertex represents a graph vertex ) 
type Vertex struct { 
	key int  // it should be holding a key 
	adjacent []*Vertex // This is a slice of pointers to vertices . 
	// Because we are implementing an adjacency list, each vertex in the graph has a list of vertices that points to the vertices of the  neighbor . 
}




//  AddNode (AddVertex) operation , we  add a Vertex to the graph  
func (g *Graph) AddVertex(k int) {
     if contains(g.vertices, k ) { 
				err := fmt.Errorf("Vertex %v not added beacause it is an existing key", k )
				fmt.Println(err.Error()) 
		   return 
		 

		}
// Creating a Vertex that has k as the key. 
			g.vertices = append(g.vertices, &Vertex{key:k}) // g.vertices is the vertices 's list . 
		 }



// Add Edge adds an edge to the graph . 
func (g *Graph) AddEdge(from, to int ) {  
	  // get address of the vertex  based on the keys.
		fromVertex := g.getVertex(from)
		toVertex := g.getVertex(to) 
		// Check  for errors . 
		if fromVertex == nil || toVertex == nil { 
			fmt.Println("Invalid vertex keys.")
			return 
		

		}
		// add edge to the adjacency list . 
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex) 


}


// getVertex returns a pointer to the Vertex with a key . 
func (g *Graph) getVertex(k int) *Vertex { 
	for i, v := range g.vertices { 
		if v.key == k { 
			return g.vertices[i] 
		}
	}

 return nil 
}


  
// contains checks if a vertex with the given key exists in the slice . 
func contains(s []*Vertex, k int) bool  {  
		for _, v := range s { 
			if k == v.key { 
				return true 
			}
		}
		return false 

}




// Print will print the adjacent list for each vertex of the graph (it will print the content of the graph) . 
func(g *Graph) Print() { 
	// the `for` loop will loop over  the vertices . 
		for _, v := range g.vertices { 
		// Printing each  keys for each vertex. 
				fmt.Printf("\n Vertex %v : ", v.key )
	// This loop iterates over the `adjacent ` slice of vertices for a specific vertex `v `
  // prints  the keys of each adjacent vertex inside the adjacency  list .
    for _, v := range v.adjacent { 
			fmt.Printf(" %v ", v.key)
		}

	}
	fmt.Println()
}





func main() {
	graph := &Graph{} 

	// We use a   `for`  loop to create nodes.
	for i := 0; i < 5; i++ { 
		graph.AddVertex(i)
	}
	fmt.Println("The list of addresses are :"  , graph, ", displayig the memory addresses of the vertex objects. 	\n Each address represents a pointer to a Vertex struct in memory.") 

	 
	 graph.AddEdge(0, 1)
	 graph.AddEdge(0, 2)
	 graph.AddEdge(1, 2)
	 graph.AddEdge(2, 0)
	 graph.AddEdge(2, 1)
	 graph.AddEdge(2, 3) 
	 graph.AddEdge(3, 1 ) 
	 graph.AddEdge(3, 0 )
	 graph.AddEdge(3, 2)
	 
	 graph.Print()


}