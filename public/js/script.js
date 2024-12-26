
/*accordion*/
$(document).ready(function () {
    $('.accord-bt').click(function () {
        var parent = $(this).closest('.sk-doc-tp');
        $(this).toggleClass('active');
        parent.find('.doc-lst').toggleClass('hidden');
        parent.find('.active').removeClass('hidden');
        parent.find('.arrow').toggleClass(' rotate-180');
    });

    // $('.hvr-bg').on('click', function () {
    //     var isVisible = $(this).next().hasClass('hidden');

    //     $('.elemt-pt').addClass('hidden').removeClass('animate-dropanime');
    //     $('.hvr-bg').removeClass('bg-[#EFEFEF]');

    //     if (isVisible) {
    //         $(this).toggleClass('bg-[#EFEFEF]');
    //         $(this).next().toggleClass('hidden animate-dropanime');
    //     }
    // });

    // $('.cls-hvr-dp').hover(function () {
    //     $('.elemt-pt').addClass('hidden');
    //     $('.elemt-pt').removeClass('animate-dropanime');
    //     $('.hvr-bg').removeClass('bg-[#EFEFEF]');
    // });

    $('.element-list a').click(function () {
        $('.elemt-pt').addClass('hidden');
        $('.elemt-pt').removeClass('animate-dropanime');
        $('.hvr-bg').removeClass('bg-[#EFEFEF]');
    });

    $(".cursor-grab").click(function() {
        $('.duplicatept').toggleClass('hidden');
    });
});

/*toggle sidebar*/
$(document).ready(function () {
    $('.toggle-button').click(function () {
        $('#sk-aside ').toggleClass('max-md:-translate-x-full');
        $('main ').toggleClass('dropbck');
    });
});

/*search*/
$(document).ready(function () {
    $(".srchBtn-togg").click(function () {
        $(this).next('input').toggleClass('w-[calc(100%-36px)] h-full block ');
        $(this).next('input').toggleClass('hidden ');
        $(this).parent('div').toggleClass('w-[32px] w-[300px] justify-start p-2.5 border border-[#ECECEC] rounded-sm gap-3 overflow-hidden ');
    });
});

/*create entires right side tab open*/

$(document).ready(function () {
    $('.tab-togg').click(function () {
        $('.editor-tabs').toggleClass('translate-x-[100%]');
        $('.editor-bdy').toggleClass('mr-[387px]');
    });
});

/*block hide and show*/
$(document).ready(function () {
    $('.hd-crd-btn').click(function () {
        $('#hd-crd, .hd-crd-btn').toggleClass('hide');
    });
});

/*07-10-24*/

$("#view-more").on("click", function () {
    $("#viewmore-content").toggleClass("hidden");
    $("#view-more").toggleClass("active");
});

document.addEventListener('DOMContentLoaded', function () {
    var tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'));
    tooltipTriggerList.forEach(function (tooltipTriggerEl) {
        new bootstrap.Tooltip(tooltipTriggerEl);
    });
});

// hover-tab

function hovertab(evt, tabid) {

   
    if (tabid=="home2"){
       
        $(".homehead").addClass("before:w-[66%]");
        $(".homehead").removeClass("before:w-[33%]");
        $(".homehead").removeClass("before:w-[100%]");
        $("#home-tab2").addClass("active")
        $("#home-tab3").removeClass("active")
    }else if(tabid=="home"){  
        $(".homehead").removeClass("before:w-[66%]");
        $(".homehead").addClass("before:w-[33%]");
        $(".homehead").removeClass("before:w-[100%]");
        $("#home-tab3").removeClass("active")
        $("#home-tab2").removeClass("active")
    }else if (tabid =="home3"){

       
        $(".homehead").addClass("before:w-[100%]");
        $(".homehead").removeClass("before:w-[66%]");
        $(".homehead").removeClass("before:w-[33%]");
        $("#home-tab3").addClass("active")

    }
    // var i, tabcontent, tablinks;
    // tabcontent = document.getElementsByClassName("tab-pane");
    // for (i = 0; i < tabcontent.length; i++) {
    //     tabcontent[i].style.display = "none";
    // }
    // tablinks = document.getElementsByClassName("nav-link-tab");
    // for (i = 0; i < tablinks.length; i++) {
    //     tablinks[i].className = tablinks[i].className.replace(" active", "");
    // }
    // document.getElementById(tabid).style.display = "block";
    // if (evt.currentTarget.classList.contains('active')) {
    //     evt.currentTarget.classList.remove('active'); // Remove active class if it exists
    // } else {
    //     evt.currentTarget.classList.add('active'); // Add active class if it doesn't exist
    // }
 }

// 
$("#home").on("click", function(){
    console.log("hjl");
    $("#home-tab").addClass("active")
})

$("#nextbutton").on("click",function(){
    console.log("happy");
})


$(document).on('click','.plussim',function(){

    if ($('.editorsec').hasClass('hidden')){

        $('.editorsec').removeClass('hidden')
    }else{
        $('.editorsec').addClass('hidden')
    }
})
// tab-Selection


const tabPanes = document.querySelectorAll('.tab-pane');
const prevButton = document.getElementById('prev');
const nextButton = document.getElementById('next');

let currentIndex = 0;

prevButton.addEventListener('click', () => {
  if (currentIndex > 0) {
    currentIndex--;
  } else {
    currentIndex = tabPanes.length - 1;
  }
  updateTabs();
});

nextButton.addEventListener('click', () => {
  if (currentIndex < tabPanes.length - 1) {
    currentIndex++;
    if (currentIndex == 1){
        $("#home-tab2").addClass("active")
        $("#home-tab3").removeClass("active")
        $("#home-tab4").removeClass("active")
        $("#home-tab").removeClass("active")
      
    }else  if (currentIndex == 2){
        $("#home-tab3").addClass("active")
        $("#home-tab2").removeClass("active")
        $("#home-tab4").removeClass("active")
        $("#home-tab").removeClass("active")
    
    }else  if (currentIndex == 3){
        $("#home-tab4").addClass("active")
        $("#home-tab2").removeClass("active")
        $("#home-tab3").removeClass("active")
        $("#home-tab").removeClass("active")
    }else  {
        $("#home-tab").addClass("active")
        $("#home-tab2").removeClass("active")
        $("#home-tab4").removeClass("active")
        $("#home-tab3").removeClass("active")
    }
  } else {
    currentIndex = 0;
  }
  updateTabs();
});

function updateTabs() {
  tabPanes.forEach((tabPane, index) => {
    if (index === currentIndex) {
        console.log("index",index);
      tabPane.style.display = "block";
      if (index == 0){
        $("#home-tab").addClass("active")
        $("#home-tab2").removeClass("active")
        $("#home-tab4").removeClass("active")
        $("#home-tab3").removeClass("active")
        
    }else  if (index == 1){
        $("#home-tab2").addClass("active")
        $("#home-tab3").removeClass("active")
        $("#home-tab4").removeClass("active")
        $("#home-tab").removeClass("active")
        
    }else  if (index == 2){
        $("#home-tab3").addClass("active")
        $("#home-tab2").removeClass("active")
        $("#home-tab4").removeClass("active")
        $("#home-tab").removeClass("active")
      
       
    }else if (index == 3){
        $("#home-tab4").addClass("active")
        $("#home-tab2").removeClass("active")
        $("#home-tab3").removeClass("active")
        $("#home-tab").removeClass("active")
    }
    } else {
      tabPane.style.display = "none";
    }
  });
}

updateTabs(); // initialize the first tab


$(document).on("mouseenter", '#home-tab2', function() {
    console.log("hover");
    $(".homehead").addClass("before:w-[66%]");
    $(".homehead").removeClass("before:w-[33%]");
}).on("mouseleave", '#home-tab2', function() {
    // Optionally, you can add functionality for when the mouse leaves
    $(".homehead").removeClass("before:w-[66%]");
    $(".homehead").addClass("before:w-[33%]");
});

