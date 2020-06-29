// Client side

$(function(){
    // WebSocket이 지원되는 브라우저 인지 체크
    if (!window.WebSocket) {
        alert("No WebSocket!")
        return
    }

    var username;
    var $chatlog = $('#chat-log')
    var $chatmsg = $('#chat-msg')

    addmessage = function(data) {
        const obj = JSON.parse(data)
        console.log("username: " + username + ", obj.username: " + obj.id)
        if (username === obj.id) {
            $chatlog.prepend("<div><span><P align=right>" + obj.data + "</P></span></div>");
        } else {
            $chatlog.prepend("<div><span>" + obj.data + "</span?</div>");
        }
    };
    
    // WebSocket에 연결
    connect = function() {
        ws = new WebSocket("ws://" + window.location.host + "/ws"); // ws://: http protocol이 아닌 web socket protocol을 사용함
                                                                    // window.location.host == "localhost",
                                                                    // but domain이 바뀔 경우 이 값이 변경될 수 있으니 window.location.host 사용
        ws.onopen = function(e) {
            console.log("onopen", arguments);
        };
        ws.onclose = function(e) {
            console.log("onclose", arguments);
        };
        ws.onmessage = function(e) {
            console.log("onmessage", arguments);
            addmessage(e.data);
        };
    }

    connect();

    var isBlank = function(string) {
        return string == null || string.trim() === "";
    };

    while (isBlank(username)) {
        username = prompt("What's your name?");
        if (!isBlank(username)) {
            $('#user-name').html('<b>' + username + '</b>');
        }
    }

    $('#input-form').on('submit', function(e) {
        if (ws.readyState === ws.OPEN) {
            ws.send(JSON.stringify({
                type: "msg",
                id: username,
                data: $chatmsg.val()
            })); // stringify : JavaScript object를 String으로 변환하여 전달
        } else {
            console.log("ws.readyState: ", ws.readyState);
        }
        $chatmsg.val("");
        $chatmsg.focus();
        return false;
    });
})