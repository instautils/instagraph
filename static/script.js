var colors = {
    'nodes': {
        'light': 'rgba(252, 227, 236, 0.05)',
        'dark': '#ad1457'
    },
    'edges': {
        'light': 'rgba(252, 227, 236, 0.05)',
        'dark': '#f06292'
    }
}
sigma.classes.graph.addMethod('neighbors', function (nodeId) {
    var k,
        neighbors = {},
        index = this.allNeighborsIndex[nodeId] || {};

    for (k in index)
        neighbors[k] = this.nodesIndex[k];

    return neighbors;
});

var s = new sigma({
    settings: {
        labelThreshold: 15,
    },
    renderer: {
        container: document.getElementById('container'),
        type: 'canvas'
    }
});

xhttp = new XMLHttpRequest();
xhttp.onreadystatechange = function () {
    if (this.readyState != 4 || this.status != 200)
        return;

    var nodes = {};
    var json = JSON.parse(this.responseText);
    var averageSize = 0;
    json.nodes.forEach(function (node) {
        averageSize += node.size;
    });
    averageSize /= json.nodes.length;
    json.nodes.forEach(function (node) {
        if (node.size < averageSize) {
            return;
        }
        nodes[node.id] = true;
        node.color = colors.nodes.dark;
        node.x *= 2;
        node.y *= 2;
        node.size *= 2;
        s.graph.addNode(node);
    });
    json.edges.forEach(function (edge) {
        if (!nodes[edge.source] || !nodes[edge.target]) {
            return;
        }
        edge.color = colors.edges.dark;
        s.graph.addEdge(edge);
    });
    s.refresh();
};
xhttp.open("GET", '/static/data.json', true);
xhttp.send();

s.bind('clickNode', function (e) {
    var nodeId = e.data.node.id,
        toKeep = s.graph.neighbors(nodeId);
    toKeep[nodeId] = e.data.node;

    s.graph.nodes().forEach(function (n) {
        if (toKeep[n.id])
            n.color = colors.nodes.dark;
        else
            n.color = colors.nodes.light;
    });

    s.graph.edges().forEach(function (e) {
        if (toKeep[e.source] && toKeep[e.target])
            e.color = colors.edges.dark;
        else
            e.color = colors.edges.light;
    });

    s.refresh();
});

s.bind('clickStage', function (e) {
    s.graph.nodes().forEach(function (n) {
        n.color = colors.nodes.dark;
    });

    s.graph.edges().forEach(function (e) {
        e.color = colors.edges.dark;
    });
    s.refresh();
});