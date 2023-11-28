# Definition
* Graph: 선으로 연결된 점들의 집합
* Matching: 짝의 집합 
    * Exposed: Matching에 못 들어간 Node
* Maximum matching: Matching들 중 edge 수 최대

## Micali, Silvio and Vazirani's algorithm
* Free vertex: 연결 안된 node
* Alternating path: 매칭O, 매칭X의 반복
* Augmenting path: 두 free vertices의 alternating path

# Idea of Edmonds' algorithm
* Augmenting paths
    * A path의 양 끝이 exposed일 때, match 상태 뒤집기
    * Edge가 +1되는 효과
    * 더 이상 없을 때까지 반복
* Blossom
    * Odd cycle, alternating edges
    * Stem: Blossom에서 연결 안 된 node 

# Algorithm
1. Exposed vertex에서 BFS
    * 처음에는 empty이므로 random하게 선택
2. Blossom을 만났다면
    1. blossom을 하나의 node로 contract
    2. 이 상태에서 Augmenting path 찾기
    3. 찾은 path를 통해 improve matching
    4. blossom을 원래대로 펼쳐 놓기 

# References
https://youtu.be/3roPs1Bvg1Q?si=2xkWkb70OkJ_wzDo