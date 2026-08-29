var languagedata

/** */
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })

    $('.Content').addClass('checked');

})



$(document).on('click', '.custom-page', function () {

    $('.pagetype').val('static-page')
    $('.pagesDropdown').removeClass('hidden')
})

$(document).on('click', '#categoryclick', function () {
    

    customepage = $(this).attr('data-name').trim()

    $('.selectpage').val(customepage)

    $('.customchooseerr').addClass('hidden')

    $(this).closest('.dropdown').find('.pagepath').val(customepage);

     $(this).closest('.dropdown-menu').removeClass('show')
    

})

$(document).on('click', '.dynamicradio', function () {

    $('.pagetype').val('editor-page')

     $('.slugdiv').removeClass('hidden')

      $('.staticpagedropdown').addClass('hidden')

    $('.landingpagedropdown').addClass('hidden')

   

})


$(document).on('click', '.blockradio', function () {

    $('.pagetype').val('block-page')

    $('.slugdiv').addClass('hidden')

      $('.staticpagedropdown').addClass('hidden')

    $('.landingpagedropdown').addClass('hidden')

   

})
$(document).on('click', '.staticradio', function () {

    $('.pagetype').val('static-page')

      $('.staticpagedropdown').removeClass('hidden')

    $('.landingpagedropdown').addClass('hidden')

    $('.selectpage').val("")

     $('.customchooseerr').addClass('hidden')

      $('.customselectpage').val("")

       $('.slugdiv').removeClass('hidden')
})

$(document).on('click', '.landingradio', function () {

    $('.pagetype').val('landing-page')

    $('.staticpagedropdown').addClass('hidden')

    $('.landingpagedropdown').removeClass('hidden')

    $('.selectpage').val("")

      $('.customchooseerr').addClass('hidden')

      $('.landingselectpage').val("")

       $('.slugdiv').removeClass('hidden')

})
$(document).on('keyup', '.customepage_name', function () {

    $('.customnameerr').addClass('hidden')

  

})

$(document).on('keyup', '.customepage_slug', function () {

    $('.customslugerr').addClass('hidden')

})
$(document).on('keyup', '.dynamicpage_name', function () {

    $('.dynamicnameerr').addClass('hidden')

})
$(document).on('keyup', '.dynamicpage_slug', function () {

    $('.dynamicslugerr').addClass('hidden')

})
$(document).on('click', '#pagemodalsave-btn', function () {

    custompagename = $('.customepage_name').val()
    pagetype = $('.pagetype').val().trim()
    selectpage = $('.selectpage').val()
    pageid = $('#page_id').val()
    var pagepath
    var flag = true
    var custompageslug
    if ((pagetype == 'static-page') ||  (pagetype == 'landing-page')) {


        pagepath = $('.selectpage').val().trim()
         


        if (selectpage == "") {

            flag = false

            $('.customchooseerr').removeClass('hidden')
        } else {

            $('.customchooseerr').addClass('hidden')
        }

    }
    templateid = $('.templateid').val()
    console.log(pagepath, pagetype, "valuess")

    custompagename = $('.dynamicpage_name').val()
        custompageslug =$('.dynamicpage_slug').val()

 
    if ((custompagename == "") ) {


        $('.dynamicnameerr').text('*Please Enter Page name').removeClass('hidden')
    }
      
    if (pagetype !="block-page"){
     
    if ((custompageslug == "") ) {

        $('.dynamicslugerr').text('*Please Enter Page Slug').removeClass('hidden')
    }
}
    isDuplicate = PageDuplicate(custompagename, pageid)

  
    
    if ((!isDuplicate) ) {

        $('.dynamicnameerr').text('*Page Name Already Exists').removeClass('hidden')
    }
     if (pagetype !="block-page"){
     
    SlugDuplicate =PageSlugDuplicate(custompageslug,pageid)

    if ((!SlugDuplicate) ) {

        $('.dynamicslugerr').text('*Slug Name Already Exists').removeClass('hidden')
    }
     }
  
    metatitle =$('#metatitle').val()
    metadesc =$('#metadesc').val()
    metakeywords =$('#metakey').val()
    metaslug =$('#metaslug').val()

    

 if (custompagename != "" && flag && isDuplicate && 
    (pagetype == "block-page" || (custompageslug != "" && SlugDuplicate))) {
        $.ajax({
            url: "/admin/website/pages/savepagedata",
            type: "POST",
            async: false,
            data: { "pageid": pageid, "page_name": custompagename,"page_slug":custompageslug, "pagetype": pagetype, "pagepath": pagepath, "webid": templateid, csrf: $("input[name='csrf']").val(),"meta_title":metatitle,"meta_description":metadesc,"meta_keywords":metakeywords,"meta_slug":metaslug },
            datatype: "json",
            caches: false,
            success: function (data) {
                console.log(data, data.page.Name, "data");
                if ((pagetype == 'static-page') ||(pagetype=="landing-page")) {

                    window.location.href = "/admin/website/pages/"
                } else if ((pagetype == 'editor-page') ||(pagetype='block-page')){
                    window.location.href = "/admin/website/pages/editpage/" + data.page.Id
                }


            }
        })
    }
})

$(document).on('click', '.pagecancelbtn', function () {

  
    $('.dynamicpage_name').val("")
     $('.dynamicpage_slug').val("")
    $('.selectpage').val("")
    $('#page_id').val("")
    $('.dynamicnameerr').addClass('hidden')
    $('.dynamicslugerr').addClass('hidden')
  
    

})
function PageDuplicate(pagename, pageid) {

    var isDuplicate = true;
    $.ajax({
        url: "/admin/website/pages/checkpagename",
        type: "POST",
        async: false,
        data: {
            "pageid": pageid,
            "page_name": pagename,
            "webid": templateid = $('.templateid').val(),
            csrf: $("input[name='csrf']").val()
        },
        datatype: "json",
        cache: false,
        success: function (data) {
            var result = data.trim();

            console.log(result, "resulttt")
            if (result == 'true') {
                isDuplicate = false;
            }
        }
    });

    return isDuplicate
}

function PageSlugDuplicate(pagename, pageid) {

    var isDuplicate = true;
    $.ajax({
        url: "/admin/checkduplicateroute",
        type: "POST",
        async: false,
        data: {
            "product_id": pageid,
            "slug_name": pagename,
            "module_name": "Pages",
            csrf: $("input[name='csrf']").val()
        },
        datatype: "json",
        cache: false,
        success: function (data) {
            var result = data.trim();

            console.log(result, "resulttt")
            if (result == 'true') {
                isDuplicate = false;
            }
        }
    });

    return isDuplicate
}

$(document).on('click','.createpage',function(){

      $('.pagetitle').text('Add New Page')
    $('#pagemodalsave-btn').text('Add Page')
})
$(document).on('click', '.createpagebtn', function () {

    $('.pagetitle').text('Add New Page')
    $('#pagemodalsave-btn').text('Add Page')

    pagename = $(this).siblings('input').val().trim()

    templateid = $('.templateid').val()

    var currentPageId = $(this).data('currentpageid');
     metatitle =$('#metatitle').val()
    metadesc =$('#metadesc').val()
    metakeywords =$('#metakey').val()
    metaslug =$('#metaslug').val()

    if (pagename == "") {

        $(this).siblings('.nameerr').text('Please Enter Page name').removeClass('hidden')
    } else if (pageNameExists(pagename, currentPageId)) {
        $(this).siblings('.nameerr').text('Page name already exists').removeClass('hidden');
    } else {
        $.ajax({
            url: "/admin/website/pages/savepagedata",
            type: "POST",
            async: false,
            data: { "page_name": pagename, "webid": templateid, csrf: $("input[name='csrf']").val(),"meta_title":metatitle,"meta_description":metadesc,"meta_keywords":metakeywords,"meta_slug":metaslug  },
            datatype: "json",
            caches: false,
            success: function (data) {
                console.log(data, data.page.Name, "data");

                window.location.href = "/admin/website/pages/editpage/" + data.page.Id 

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


$(document).on('click', '.editpagename', function () {

  $('.pagetitle').text('Edit Page')
    $('#pagemodalsave-btn').text('Update Page')
    pageid = $(this).attr('data-id')
    pagename = $(this).attr('data-pagename')
    pagepath = $(this).attr('data-path')
    pagetype = $(this).attr('data-type').trim()
    pageslug =$(this).attr('data-slug')
    $('#page_id').val(pageid)

    
     
    if (pagetype == "static-page") {
       
        //    $('#profile-tab').addClass('active')
        //    $('#profile-tab-pane').removeClass('active')
     
        $('.customselectpage').val(pagepath)
        $('#profile-tab').trigger('click')
        $('#radio1').prop('checked', true).trigger('change');
        $('.staticpagedropdown').removeClass('hidden')
        $('.landingpagedropdown').addClass('hidden')

    } else if (pagetype =="landing-page"){
       
       
        $('.landingselectpage').val(pagepath)
        $('#landing-page').trigger('click')
        $('#radiolanding').prop('checked', true).trigger('change');
        $('.staticpagedropdown').addClass('hidden')
        $('.landingpagedropdown').removeClass('hidden')

    }else if (pagetype =="block-page"){
       
       
     
        $('.blockradio').trigger('click')
        $('#radioblock').prop('checked', true).trigger('change');
        $('.staticpagedropdown').addClass('hidden')
        $('.landingpagedropdown').addClass('hidden')

    } else {

        // $('#profile-tab-pane').addClass('active')
        //  $('#profile-tab').removeClass('active')
       
        $('#home-tab').trigger('click')
        $('#radio').prop('checked', true).trigger('change');
    }
    $('.pagetype').val(pagetype)
     $('.dynamicpage_name').val(pagename)
     $('.dynamicpage_slug').val(pageslug)
     $('.selectpage').val(pagepath)



    // $(this).siblings('.pagename-link').addClass('hidden');
    // $(this).siblings('div').children('.pagename-input').removeClass('hidden').focus();
    // $(this).addClass('hidden');
    // $(this).siblings('.editpagebtn').removeClass('hidden');
});

$(document).on('click', '.editpagebtn', function () {
    var newName = $(this).siblings('.newdiv').find('.pagename-input').val();
    var pageid = $(this).attr('data-id');
    var $btn = $(this);

    var currentPageId = $(this).data('currentpageid');
    $btn.siblings('.newdiv').find('.nameerr').addClass('hidden');
    templateid = $('.templateid').val()
     metatitle =$('#metatitle').val()
    metadesc =$('#metadesc').val()
    metakeywords =$('#metakey').val()
    metaslug =$('#metaslug').val()
    if (newName == "") {


        $btn.siblings('.newdiv').find('.nameerr').text('Please Enter Page Name').removeClass('hidden');
    } else if (pageNameExists(newName, currentPageId)) {
        $btn.siblings('.newdiv').find('.nameerr').text('Page name already exists').removeClass('hidden');
    } else {
        $.ajax({
            url: "/admin/website/pages/savepagedata",
            type: "POST",
            async: false,
            data: { "page_name": newName, "webid": templateid, "pageid": pageid, csrf: $("input[name='csrf']").val(),"meta_title":metatitle,"meta_description":metadesc,"meta_keywords":metakeywords,"meta_slug":metaslug  },
            datatype: "json",
            cache: false,
            success: function (data) {
                setCookie("get-toast", "Page Created Successfully")
                window.location.href = "/admin/website/pages/editpage/" + data.page.Id 
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
    $('.pagename-link').each(function () {
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
    $('.pagename-link').each(function () {

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

    console.log(spurtdata, "spurtdataa")

    pageid = $('.pageid').val()

    templateid = $('.templateid').val()

    description = spurtdata.html
   metatitle =$('#metatitle').val()
    metadesc =$('#metadesc').val()
    metakeywords =$('#metakey').val()
    metaslug =$('#metaslug').val()
  

    console.log("checkvaluess",metatitle,metadesc)

    $.ajax({
        url: "/admin/website/pages/savepagedata",
        type: "POST",
        async: false,
        data: { "pageid": pageid, "webid": templateid, "pagedata": description, csrf: $("input[name='csrf']").val(),"meta_title":metatitle,"meta_description":metadesc,"meta_keywords":metakeywords,"page_slug":metaslug  },
        datatype: "json",
        cache: false,
        success: function (data) {

            setCookie("get-toast", "Page Updated Successfully")
             window.location.href = "/admin/website/pages/" 

        }
    })

}
)
$(document).on('click', '#pagesave', function () {


    url =window.location.href

    if (url.includes('/createpage')){

        console.log("checkonee")

         notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li> <div class="toast-msg flex  max-sm:max-w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify" > <img src="/public/img/close-toast.svg" alt="close"> </a> <div> <img src="/public/img/danger-group-12.svg" alt="toast error"> </div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Warning</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] ">Please Choose Page</p></div></div> </li></ul> `;
                    $(notify_content).insertBefore(".header-rht");
                    setTimeout(function () {
                        $('.toast-msg').fadeOut('slow', function () {
                            $(this).remove();

                        });
                    }, 5000);
    }else{


 

    const event = new CustomEvent("getHTML", {
    });
    document.dispatchEvent(event);

}
})


$(document).on('keyup', '.pagenameinput', function () {

    if ($(this).val() != "") {
        $(this).siblings('.nameerr').addClass("hidden")
    } else {
        $(this).siblings('.nameerr').removeClass("hidden")

    }

})

$("#searchcatlists").keyup(function () {
    var found = false;
    var searchTerm = $(this).val().trim().toLowerCase()


    console.log("isvisible", searchTerm)
    $(".page-dropdownlist ").filter(function () {
        var isVisible = $(this).attr('data-name').toLowerCase().indexOf(searchTerm) > -1;

        console.log("isvisible", isVisible)
        $(this).toggle(isVisible);
        if (isVisible) found = true;
    })
    if (found) {
        $('.pagelistnodata').addClass('hidden');
    } else {
        console.log("checkconditon")
        $('.pagelistnodata').removeClass('hidden');
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

 function toggleAccordionindustrydetail(index) {
    console.log("lkjkjk");
    
    const content = document.getElementById(`content-${index}`);
    const button = document.getElementById(`btn-${index}`);
    const icon = document.getElementById(`icon-${index}`);
    const wrapper = document.querySelector(`.accord-${index}`);

    const minusSVG = `<img src="/picco_template/assets/img/accord-up.svg" alt="">`;
    const plusSVG = `<img src="/picco_template/assets/img/accord-down.svg" alt="">`;

    if (content.style.maxHeight && content.style.maxHeight !== '0px') {
        content.style.maxHeight = '0';
        icon.innerHTML = plusSVG;
        button.classList.remove('active');
        wrapper.classList.remove('active')
    } else {
        content.style.maxHeight = content.scrollHeight + 'px';
        icon.innerHTML = minusSVG;
        button.classList.add('active');
        wrapper.classList.add('active')
    }
}
// 

// JAN 13

function toggleAccordionfaq(index) {
    const content = document.getElementById(`content-${index}`);
    const button = document.getElementById(`btn-${index}`);
    const icon = document.getElementById(`icon-${index}`);
    const wrapper = document.querySelector(`.accord-${index}`);
  
    const minusSVG = `<img src="/picco_template/assets/img/accord-minus-dark.svg" alt="">`;
    const plusSVG = `<img src="/picco_template/assets/img/accord-plus-dark.svg" alt="">`;
  
    if (content.style.maxHeight && content.style.maxHeight !== '0px') {
      content.style.maxHeight = '0';
      icon.innerHTML = plusSVG;
      button.classList.remove('active');
      wrapper.classList.remove('active')
    } else {
      content.style.maxHeight = content.scrollHeight + 'px';
      icon.innerHTML = minusSVG;
      button.classList.add('active');
      wrapper.classList.add('active')
    }
  }


 

async function openItem(item) {
    if (!item) return;

    const desc = item.querySelector('.desc');
    if (!desc) return;
console.log(desc,"desc");

    item.classList.add('industries-titles-right-active');
    desc.classList.add('open-desc');
console.log(desc,"desc", desc.classList);

    await animateHeight(desc, 0, desc.scrollHeight, OPEN_TIME);
}


//   

async function closeItem(item) {
    if (!item) return;

    const desc = item.querySelector('.desc');
    if (!desc) return;

    await animateHeight(desc, desc.scrollHeight, 0, CLOSE_TIME);

    desc.classList.remove('open-desc');
    item.classList.remove('industries-titles-right-active');
}

// 

const OPEN_TIME = 400;
const CLOSE_TIME = 300;

function animateHeight(el, from, to, duration) {
  return new Promise(resolve => {
    el.style.overflow = 'hidden';
    el.style.maxHeight = from + 'px';
    el.style.transition = `max-height ${duration}ms ease`;

    requestAnimationFrame(() => {
      el.style.maxHeight = to + 'px';
    });

    setTimeout(() => {
      if (to !== 0) el.style.maxHeight = 'none';
      resolve();
    }, duration);
  });
}

// 
  

  $(document).on('click', '.industries_titles_right', async function () {
    console.log('clicked element:', this);

    await openItem(this);
    activeIndex = this;
});

// For carrer page slider

document.addEventListener("DOMContentLoaded", function () {



  const slider = document.getElementById("slider");
  if (!slider) return; 
  const slides = slider.children;
  const totalSlides = slides.length;

  const nextBtn = document.getElementById("next");
  const prevBtn = document.getElementById("prev");
  if (!nextBtn || !prevBtn) return;

  let index = 0;
  let interval;
  const slideWidth = 75; // 70% + 5% gap

  function updateSlide() {
    slider.style.transform = `translateX(-${index * slideWidth}%)`;
  }

  function nextSlide() {
    index = (index + 1) % totalSlides;   // 🔁 LOOP
    updateSlide();
  }

  function prevSlide() {
    index = (index - 1 + totalSlides) % totalSlides; // 🔁 LOOP BACK
    updateSlide();
  }
  function resetAutoSlide() {
    clearInterval(interval);
    startAutoSlide();
  }

  nextBtn.addEventListener("click", () => {
    nextSlide();
    resetAutoSlide();
  });

  prevBtn.addEventListener("click", () => {
    prevSlide();
    resetAutoSlide();
  });
});
$(function () {
  const $slider = $('#slider');
  if (!$slider.length) return;
  const $slides = $slider.children();
  const $prev = $('#prev');
  const $next = $('#next');

  let currentIndex = 0;
  const totalSlides = $slides.length;
  const animationSpeed = 1000; // adjust speed here
  let isAnimating = false;

  $slider.css({
    transition: `transform ${animationSpeed}ms ease-in-out`
  });

  function updateSlider() {
    const slideWidth = $slides.outerWidth(true);
    const translateX = -(slideWidth * currentIndex);

    isAnimating = true;
    $slider.css('transform', `translateX(${translateX}px)`);

    setTimeout(() => {
      isAnimating = false;
    }, animationSpeed);
  }

  function goNext() {
    if (isAnimating) return;
    if (currentIndex < totalSlides - 1) {
      currentIndex++;
      updateSlider();
    }
  }

  function goPrev() {
    if (isAnimating) return;
    if (currentIndex > 0) {
      currentIndex--;
      updateSlider();
    }
  }

  // Buttons
  $next.on('click', goNext);
  $prev.on('click', goPrev);

  // Keyboard
  $(document).on('keydown', function (e) {
    if (e.key === "ArrowRight") goNext();
    if (e.key === "ArrowLeft") goPrev();
  });

  // Mouse wheel (one move only)
  $('.relative.overflow-hidden').on('wheel', function (e) {
    e.preventDefault();
    if (isAnimating) return;

    if (e.originalEvent.deltaY > 0) {
      goNext();
    } else {
      goPrev();
    }
  });
});
// ===================

// service detail
// Tab switching functionality for responsive design
function switchTab(tabName) {
  // Remove active class from all tabs
  document.querySelectorAll('.tab-btn').forEach(btn => {
    btn.classList.remove('active');
    btn.classList.remove('text-[#FFFFFF]', 'bg-[#1D1D1B]', 'border-[#1D1D1B]');
    btn.classList.add('text-[#3B4045]');
  });
  
  // Hide all tab content
  document.querySelectorAll('.tab-content').forEach(content => {
    content.classList.add('hidden');
  });
  
  // Show selected tab content
  const selectedContent = document.getElementById(tabName + '-content');
  if (selectedContent) {
    selectedContent.classList.remove('hidden');
  }
  
  // Add active class to selected tab
  const selectedTab = document.getElementById(tabName + '-tab');
  if (selectedTab) {
    selectedTab.classList.add('active');
    selectedTab.classList.add('text-[#FFFFFF]', 'bg-[#1D1D1B]', 'border-[#1D1D1B]');
    selectedTab.classList.remove('text-[#3B4045]');
  }
}

// Function to generate slug from text
function generateSlug(text) {
    return text
        .toLowerCase()
        .trim()
        .replace(/[^\w\s-]/g, '')        
        .replace(/\s+/g, '-')             
        .replace(/-+/g, '-')              
        .replace(/^-+|-+$/g, '');         
}

$(document).on('keyup', '.dynamicpage_name', function () {
    let name = $(this).val();
    let slugField = $('.dynamicpage_slug');
    const specialCharRegex = /[^\w\s-]/g;
    
    if (specialCharRegex.test(name)) {
        $(this).siblings('.dynamicnameerr').text('Special characters are not allowed').removeClass('hidden');
    } else {
        $(this).siblings('.dynamicnameerr').addClass('hidden');
    }

    // Always generate slug regardless of special char check
    if (!slugField.data('manual')) {
        let slug = generateSlug(name);
        slugField.val(slug);
    }
});

$(document).on('keypress', '.dynamicpage_name', function (e) {
    const char = String.fromCharCode(e.which);
    if (!/^[a-zA-Z0-9\-_ ]$/.test(char)) {
        e.preventDefault();
        $('.error-dynamicname').remove();
        const $errorMessage = $('<div class="error-dynamicname text-xs mt-1" style="color: #F26674;">* Only letters, numbers, - and _ are allowed</div>');
        $(this).after($errorMessage);
    } else {
        $('.error-dynamicname').remove();
    }
});

$(document).on('keypress', '.dynamicpage_slug', function (e) {
    const char = String.fromCharCode(e.which);
    if (!/^[a-zA-Z0-9\-_]$/.test(char)) {
        e.preventDefault();
        $('.error-dynamicslug').remove();
        const $errorMessage = $('<div class="error-dynamicslug text-xs mt-1" style="color: #F26674;">* Only letters, numbers, - and _ are allowed</div>');
        $(this).after($errorMessage);
        $(this).data('invalid-pressed', true);
    } else {
        $(this).data('invalid-pressed', false);
        $('.error-dynamicslug').remove();
    }
});

$(document).on('keyup', '.dynamicpage_slug', function () {
    $('.dynamicslugerr').addClass('hidden');
    const val = $(this).val();
    const specialCharRegex = /[^\w\-]/g;

    if (specialCharRegex.test(val)) {
        $(this).val(val.replace(/[^\w\-]/g, ''));
        $('.error-dynamicslug').remove();
        const $errorMessage = $('<div class="error-dynamicslug text-xs mt-1" style="color: #F26674;">* Only letters, numbers, - and _ are allowed</div>');
        $(this).after($errorMessage);
        $(this).data('invalid-pressed', false);
    } else if ($(this).data('invalid-pressed')) {
    } else {
        $('.error-dynamicslug').remove();
    }
    if (val.trim() === '') {
        $(this).data('manual', false);
    } else {
        $(this).data('manual', true);
    }
});

// AFTER:
const carouselTrack = document.querySelector('#carousel-track');
if (carouselTrack) { // ← guard
  const carouselWrapper = carouselTrack.closest('.overflow-hidden');
  if (carouselWrapper) { // ← guard the wrapper too
    carouselWrapper.style.overflowY = 'visible';
    carouselWrapper.addEventListener('wheel', function(e) {
      if (!e.shiftKey) {
        window.scrollBy(0, e.deltaY);
        e.preventDefault();
      }
    }, { passive: false });
  }
}


function fixLinkRowScroll() {
  
  const linkRows = document.querySelectorAll('a.link-row, button.link-row');
  
  linkRows.forEach(linkRow => {
    
    if (linkRow.dataset.scrollFixed) return;
    linkRow.dataset.scrollFixed = 'true';
    
    const originalOverflow = linkRow.style.overflow;
    
    
    linkRow.addEventListener('mouseenter', function() {
      this.style.overflow = 'visible';
      this.style.pointerEvents = 'auto';
      this.classList.add('hover-active');
    });
    
    linkRow.addEventListener('mouseleave', function() {
      this.style.overflow = originalOverflow || 'hidden';
      this.style.pointerEvents = 'auto';
      this.classList.remove('hover-active');
    });
    
    
    linkRow.addEventListener('wheel', function(e) {
      e.stopPropagation();
  
      window.scrollBy(0, e.deltaY * 2);
    }, { passive: false });
    
    linkRow.addEventListener('touchmove', function(e) {
      e.stopPropagation();
    }, { passive: true });
  });
}


fixLinkRowScroll();


setInterval(fixLinkRowScroll, 500);

const observer = new MutationObserver(fixLinkRowScroll);
observer.observe(document.body, { 
  childList: true, 
  subtree: true 
});

document.addEventListener('wheel', function(e) {
  const hoveredLink = e.target.closest('a.link-row.hover-active, button.link-row.hover-active');
  if (hoveredLink) {
        return;
  }
}, { passive: true });