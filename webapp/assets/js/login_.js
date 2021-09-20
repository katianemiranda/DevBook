$('#login').on('submit', fazerLogin);

function fazerLogin(evento) {
    evento.preventDefault()

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            senha: $('#senha').val(),
        }
    }).done(function() {
        console.log("aqui1");
        window.location = "/home";
    }).fail(function() {
        console.log("aqui2");
       // alert("Usuário ou senha inválidos kkkk")
    });
}