require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap-sass/assets/javascripts/bootstrap.js");

$('.nav-sidebar li').on('click', function () {
    $(this).addClass('active');
})
