var colors = {
    'light': 'rgba(252, 227, 236, 0.05)',
    'dark': '#f06292',
    'blue': '#303f9f',
}

var global = {
    'username': 'aidenzibaei',
    'min_size': 1,
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

function findNode(node) {
    console.log(node);

    var nodeId = node.id;
    var toKeep = s.graph.neighbors(nodeId);
    toKeep[nodeId] = node;

    s.graph.nodes().forEach(function (n) {
        if (toKeep[n.id])
            n.color = colors.dark;
        else
            n.color = colors.light;
    });

    s.graph.edges().forEach(function (e) {
        if (toKeep[e.source] && toKeep[e.target])
            e.color = colors.dark;
        else
            e.color = colors.light;
    });

    $(".button-untrack").removeAttr('disabled');
    s.refresh();
}

s.bind('clickNode', function (e) {
    var node = e.data.node;
    findNode(node);
});

function fetchAndInit() {
    var deferred = $.Deferred();
    $.get({
        url: '/static/data.json'
    }).then(function (json) {
        var nodes = {};
        json.nodes.forEach(function (node) {
            if (node.size < global.min_size) {
                return;
            }
            var factor = node.size > 10 ? 10 : node.size;
            factor = 15 - factor;
            node.size *= factor;
            node.x *= factor;
            node.y *= factor;
            nodes[node.id] = true;
            node.color = colors.dark;
            s.graph.addNode(node);
        });
        json.edges.forEach(function (edge) {
            if (!nodes[edge.source] || !nodes[edge.target]) {
                return;
            }
            edge.color = colors.dark;
            s.graph.addEdge(edge);
        });
        s.refresh();
        deferred.resolve();
    });
    return deferred.promise();
}

$(document).ready(function () {
    $('.button-untrack').attr('disabled', 'disabled');
    fetchAndInit().then(function () {
        $(".button-untrack").click(function () {
            s.graph.nodes().forEach(function (n) {
                n.color = colors.dark;
            });
            s.graph.edges().forEach(function (e) {
                e.color = colors.dark;
            });
            s.refresh();

            $(this).attr('disabled', 'disabled');
        });
    });

    $(".button-find").click(function() {
        var query = $(".gui-box input.search-node").val();
        if (query.length < 3) {
            return;
        }        
        var nodes = s.graph.nodes();
        for (var i=0;i<nodes.length;i++) {
            var node = nodes[i];
            if (node.id == query) {
                findNode(node);
                return;
            }
        }
    });

    $(".button-path").click(function() {

    });
})

