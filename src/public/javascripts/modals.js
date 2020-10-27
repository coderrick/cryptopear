var x = 0;

$('.addNewImage').click(function(){
  $('#AddImageModal').addClass('is-active');
});

$('.delete').click(function(){
  $('.modal').removeClass('is-active');
});

$('.modal-background').click(function(){
  $('.modal').removeClass('is-active');
});

$('.cancel').click(function(){
  $('.modal').removeClass('is-active');
});

$('#CompleteAddImage').click(function(){
  x++;
  $('.addNewImage').before("<div class='image'><img src='https://i.ebayimg.com/images/g/gCoAAOSwGYZdtQcM/s-l640.jpg'/><div class='desc'><div class='desc'>iPhone X 256</div><div class='desc'>$988.00</div><button class='mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--accent' style='color:white;'>Buy</button>"+$('#exampleDesc').val()+"</div><div class='dropdown'><button href='#' class='dropbtn'><i class='fas fa-wrench'></i></button><div class='dropdown-content'><a class='edit' href='#'><input hidden value='"+x.toString()+"' name='id'></input>Editar</a><a class='deleteImage' href='#'>Remover</a></div></div></div>");
  $('#exampleDesc').val('');
  $('.modal').removeClass('is-active');
  })

$(document).on('click','.deleteImage',function(){
  $(this).parent().parent().parent().remove();
});

var editing = 0;

$(document).on('click','.edit',function(){
  editing = $(this).find('input[name="id"]').val();
  $('#exampleEditDesc').val($(this).parent().parent().parent().find('.desc').text());
  $('#EditImageModal').addClass('is-active');
});

$('#CompleteEditImage').click(function(){
  $(document).find('input[name="id"][value="'+editing+'"]').parent().parent().parent().parent().find('.desc').text($('#exampleEditDesc').val());
  $('#exampleEditDesc').val('');
  $('.modal').removeClass('is-active');
});