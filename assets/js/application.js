require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap-sass/assets/javascripts/bootstrap.js");

// $(document).ready(function(){
//
//
//
// });
$('.nav-sidebar li').on('click', function () {
    $(this).addClass('active');
})
