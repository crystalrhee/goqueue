$(document).ready(function(){
  $('a').on('click', function(event) {
    if (this.hash !== '') {
      event.preventDefault();
      var hash = this.hash;
      $('html, body').animate({
        scrollTop: $(hash).offset().top
      }, 1000, function(){
        window.location.hash = hash;
      });
    }
  });

  // hamburger menu
  $('.header .menu').on('click', function() {
    if($('.header').hasClass('open')) {
      console.log('close');
      $('.header').removeClass('open');
    } else {
      console.log('open');
      $('.header').addClass('open');
    }

    return false;
  });

  // add class to dancing gifs
  $('#join').click(function(){
    $('#dance').addClass('active');
    $('.container').addClass('active');
  });
});

