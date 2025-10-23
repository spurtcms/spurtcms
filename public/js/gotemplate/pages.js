var languagedata

/** */
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
        })

            $('.Content').addClass('checked');

})


$(document).on('click','.createpage',function(){


    $('.createpagediv').removeClass('hidden')
})


$(document).on('click','.createpagebtn',function(){

   pagename= $(this).siblings('input').val().trim()

   templateid =$('.templateid').val()

   var currentPageId = $(this).data('currentpageid');

   if (pagename==""){

    $(this).siblings('.nameerr').text('Please Enter Page name').removeClass('hidden')
   }else if(pageNameExists(pagename, currentPageId)) {
         $(this).siblings('.nameerr').text('Page name already exists').removeClass('hidden');
    } else {
       $.ajax({
               url: "/admin/website/pages/savepagedata",
               type: "POST",
               async: false,
               data: { "page_name": pagename,"webid":templateid, csrf: $("input[name='csrf']").val() },
               datatype: "json",
               caches: false,
             success: function (data) {
    console.log(data, data.page.Name, "data");

    window.location.href="/admin/website/pages/editpage/"+data.page.Id+"?webid="+data.page.WebsiteId

    // var str = `<div class="flex items-center hover:bg-[#F5F5F5]">
    //             <a href="/admin/website/pages/editpage/`+data.page.Id+`?webid=`+templateid+`" class="pagename-link text-[14px] grow font-light leading-[15px] text-[#282322] p-[16px_20px] block">` + data.page.Name + `</a>
    //             <a href="javascript:void(0);" class="editpagename shrink-0 px-3 py-[20px] hover:bg-[#e1e1e1]">
    //               <img src="/public/img/edit-event.svg" alt="" class="w-[1rem]">
    //             </a>
    //              <input type="text" class="pagename-input text-[14px] font-light leading-[18px] placeholder:text-[#757575] w-full hidden" value="` + data.page.Name + `">
    //             <a href="javascript:void(0)" data-id="` + data.page.Id + `" class="editpagebtn text-[13px] px-3 py-[20px] font-normal leading-[16px] text-[#10A37F] hidden">Save</a>
               
    //             <p class="nameerr relative bottom-0 hidden text-[#F26674] text-xs mt-1">Please Enter Page Name</p>
    //           </div>`;

    // $('.pages-container').append(str);

    // // Optional: clear input and hide input form
    // $('.createpagediv input').val('');
    // $('.createpagediv').addClass('hidden');
               }
          })
   }
})


$(document).on('click', '.editpagename', function() {

   
   
    $(this).siblings('.pagename-link').addClass('hidden');
    $(this).siblings('div').children('.pagename-input').removeClass('hidden').focus();
    $(this).addClass('hidden');
    $(this).siblings('.editpagebtn').removeClass('hidden');
});

$(document).on('click', '.editpagebtn', function() {
    var newName = $(this).siblings('.newdiv').find('.pagename-input').val();
    var pageid = $(this).attr('data-id');
    var $btn = $(this);

   var currentPageId = $(this).data('currentpageid'); 
    $btn.siblings('.newdiv').find('.nameerr').addClass('hidden');
 templateid =$('.templateid').val()
    if (newName == "") {

 
        $btn.siblings('.newdiv').find('.nameerr').text('Please Enter Page Name').removeClass('hidden');
    } else if(pageNameExists(newName, currentPageId)) {
        $btn.siblings('.newdiv').find('.nameerr').text('Page name already exists').removeClass('hidden');
    } else {
        $.ajax({
            url: "/admin/website/pages/savepagedata",
            type: "POST",
            async: false,
            data: { "page_name": newName,"webid":templateid, "pageid": pageid, csrf: $("input[name='csrf']").val() },
            datatype: "json",
            cache: false, 
            success: function(data) {
    setCookie("get-toast", "Page Created Successfully")
    window.location.href="/admin/website/pages/editpage/"+data.page.Id+"?webid="+data.page.WebsiteId
                // $btn.siblings('.pagename-link').text(newName);
                // $btn.siblings('.newdiv').find('.pagename-input').addClass('hidden');
                // $btn.addClass('hidden');
                // $btn.siblings('.editpagename').removeClass('hidden');
                // $btn.siblings('.pagename-link').removeClass('hidden');
            }
        });
    }
});

function pageNameExists(name, currentPageId = null) {
    var exists = false;
    $('.pagename-link').each(function() {
        // Get the page id of this link
        var pageId = $(this).data('pageid');
        // Skip the current editing page
        if (currentPageId && pageId == currentPageId) return true;
        if ($(this).text().trim().toLowerCase() === name.trim().toLowerCase()) {
            exists = true;
            return false; // Stop loop
        }
    });
    return exists;
}

function pageNameExists(name, currentPageId = null) {
    var exists = false;
    $('.pagename-link').each(function() {
       
        var pageId = $(this).data('pageid');
       
        if (currentPageId && pageId == currentPageId) return true;
        if ($(this).text().trim().toLowerCase() === name.trim().toLowerCase()) {
            exists = true;
            return false; 
        }
    });
    return exists;
}





 document.addEventListener('saveChange', function (event) {
    spurtdata = event.detail

    console.log(spurtdata,"spurtdataa")

    pageid =$('.pageid').val()
    
    templateid =$('.templateid').val()

    description =spurtdata.html

     $.ajax({
            url: "/admin/website/pages/savepagedata",
            type: "POST",
            async: false,
            data: {  "pageid": pageid,"webid":templateid, "pagedata":description, csrf: $("input[name='csrf']").val() },
            datatype: "json",
            cache: false,
            success: function(data) {

                 setCookie("get-toast", "Page Updated Successfully")
                window.location.href="/admin/website/pages/"+templateid

            }
        })

 }
)
$(document).on('click','#pagesave',function(){

  

       const event = new CustomEvent("getHTML", {
    });
    document.dispatchEvent(event);
})


$(document).on('keyup','.pagenameinput',function(){

   if( $(this).val()!=""){
 $(this).siblings('.nameerr').addClass("hidden")
    }else{
$(this).siblings('.nameerr').removeClass("hidden")

    }
   
})

$("#searchcatlists").keyup(function () {
  var found = false;
  var searchTerm = $(this).val().trim().toLowerCase()

  
    console.log("isvisible",searchTerm)
  $(".page-dropdownlist ").filter(function () {
    var isVisible = $(this).attr('data-name').toLowerCase().indexOf(searchTerm) > -1;

    console.log("isvisible",isVisible)
    $(this).toggle(isVisible);
    if (isVisible) found = true;
  })
  if (found) {
    $('.noData-foundWrapper').hide();
  } else {
    $('.noData-foundWrapper').show();
  }
})

$(document).ready(function () {
 
    // $(".editor-tabs").addClass("translate-x-full");
 
    $(".tabclose1").on("click", function () {
        $(".editor-tabs").toggleClass("translate-x-full");
        $(".tabclose2").removeClass("hidden");
        $(".tabclose1").addClass("hidden");
    });
 
    $(".tabclose2").on("click", function () {
        $(".editor-tabs").toggleClass("translate-x-full");
        $(".tabclose1").removeClass("hidden");
        $(".tabclose2").addClass("hidden");
    });
});