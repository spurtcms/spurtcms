/**Variables */
var languagedata;

let editor;

let ckeditor1;

let collapsecount = 1

let newPage = 0

let OrderIndex = 1

let PageId = 1

let SubPageId = 1

let GroupId = 1

/** Create array */
let newpagesarr = []

let newGrouparr = []

let newSubpagearr = []

/**delete array*/
let deletePagearr = []

let deletegrouparr = []

let deletesubarr = []

/** update array */
let updPagesarr = []

let updGrouparr = []

let updSubpagearr = []

/** Temp array */
let tempPagesarr = []

let tempGrouparr = []

let tempSubpagearr = []

let NewSubObj = []

let overallarray = []

const wpm = 150;

$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })


    CKEDITORS()

    $.ajax({
        type: "get",
        url: "/spaces/pagedata",
        dataType: 'json',
        data: {
            sid: $("#spid").val(),
        },
        cache: false,
        success: function (result) {

            console.log("data,", result)

            if (result.group != null) {

                overallarray = overallarray.concat(result.group)
            }

            if (result.pages != null) {

                overallarray = overallarray.concat(result.pages)

            }

            overallarray.sort((a, b) => {

                return a.OrderIndex - b.OrderIndex;
            })

            if (result.pages != null) {

                for (let [index, x] of result.pages.entries()) {

                    var obj = {}

                    obj.PageId = x.PgId

                    obj.NewPageId = 0

                    obj.Name = x.Name

                    obj.Content = x.Content

                    obj.PageGroupId = x.Pgroupid

                    obj.NewGroupId = 0

                    obj.OrderIndex = x.OrderIndex

                    obj.ParentId = 0

                    obj.Username = x.Username

                    obj.LastUpdate = x.LastUpdate

                    obj.ReadTime =x.ReadTime

                    obj.Log = x.Log

                    var date = get_time_diff(x.LastUpdate);

                    if (date < 60) {

                        obj.LastUpdateMinutes = date + " minutes ago"

                    } else {

                        obj.LastUpdate = x.LastUpdate;

                    }

                    tempPagesarr.push(obj)
                }

            }

            if (result.group != null) {

                for (let x of result.group) {

                    var obj = {}

                    obj.GroupId = x.GroupId

                    obj.NewGroupId = 0

                    obj.Name = x.Name

                    obj.OrderIndex = x.OrderIndex

                    tempGrouparr.push(obj)
                }

            }


            if (result.subpage != null) {

                for (let x of result.subpage) {

                    var obj = {}

                    obj.SubPageId = x.SpgId

                    obj.NewSubPageId = 0

                    obj.Content = x.Content

                    obj.PageGroupId = x.PgroupId

                    obj.OrderIndex = x.OrderIndex

                    obj.ParentId = x.ParentId

                    obj.NewParentId = x.NewParentId

                    obj.Name = x.Name

                    obj.LastUpdate = x.LastUpdate

                    obj.Log = x.Log

                    obj.ReadTime = x.ReadTime

                    tempSubpagearr.push(obj)

                }

            }

            console.log(overallarray);

            if (overallarray.length == 0) {

                var pagename = "Introduction"

                var pagedata = AddPageString(pagename, PageId, languagedata, 0, OrderIndex, true)

                $('.newpages').append(pagedata)

                var pageobj = PageObjectCreate();

                pageobj.Name = pagename

                pageobj.OrderIndex = OrderIndex

                newpagesarr.push(pageobj)

                PageId++;

                OrderIndex++;

            }

            PGList()

            PageSelect()

        }


    })





})

function CKEDITORS() {

    // var url = window.location.origin;

    var url = $("#url").val();

    CKEDITOR.ClassicEditor.create(document.getElementById("neweditor"), {
        toolbar: {
            items: ['heading', 'bold', 'italic', 'blockQuote', 'imageUpload', 'numberedList', 'bulletedList', 'horizontalLine', 'link', 'code', 'codeBlock'],
            shouldNotGroupWhenFull: true
        },
        // toolbar:{
        //     items:['alignment:left', 'alignment:right', 'alignment:center', 'alignment:justify', 'alignment', 'bold', 'italic', 'underline', 'strikethrough', 'superscript', 'subscript', 'code', 'blockQuote', 'codeBlock', 'selectAll', 'undo', 'redo', 'heading', 'imageTextAlternative', 'toggleImageCaption', 'imageStyle:inline', 'imageStyle:alignLeft', 'imageStyle:alignRight', 'imageStyle:alignCenter', 'imageStyle:alignBlockLeft', 'imageStyle:alignBlockRight', 'imageStyle:block', 'imageStyle:side', 'imageStyle:wrapText', 'imageStyle:breakText', 'imageUpload', 'insertImage', 'imageInsert', 'indent', 'outdent', 'textPartLanguage', 'link', 'linkImage', 'numberedList', 'bulletedList', 'todoList', 'mediaEmbed', 'findAndReplace', 'fontBackgroundColor', 'fontColor', 'fontFamily', 'fontSize', 'highlight:yellowMarker', 'highlight:greenMarker', 'highlight:pinkMarker', 'highlight:blueMarker', 'highlight:redPen', 'highlight:greenPen', 'removeHighlight', 'highlight', , 'htmlEmbed', 'pageBreak', 'removeFormat', 'restrictedEditingException', 'showBlocks', 'style', 'specialCharacters', 'insertTable', 'tableColumn', 'tableRow', 'mergeTableCells', 'tableCellProperties', 'tableProperties', 'toggleTableCaption', 'sourceEditing']
        // },
        list: {
            properties: {
                styles: true,
                startIndex: true,
                reversed: true
            }
        },
        heading: {
            options: [
                { model: 'paragraph', title: 'Paragraph', class: 'ck-heading_paragraph' },
                { model: 'heading1', view: 'h1', title: 'Heading 1', class: 'ck-heading_heading1' },
                { model: 'heading2', view: 'h2', title: 'Heading 2', class: 'ck-heading_heading2' },
                { model: 'heading3', view: 'h3', title: 'Heading 3', class: 'ck-heading_heading3' },
                { model: 'heading4', view: 'h4', title: 'Heading 4', class: 'ck-heading_heading4' },
                { model: 'heading5', view: 'h5', title: 'Heading 5', class: 'ck-heading_heading5' },
                { model: 'heading6', view: 'h6', title: 'Heading 6', class: 'ck-heading_heading6' }
            ]
        },
        placeholder: languagedata.Spaces.pllckeditor,
        fontFamily: {
            options: [
                'default',
                'Arial, Helvetica, sans-serif',
                'Courier New, Courier, monospace',
                'Georgia, serif',
                'Lucida Sans Unicode, Lucida Grande, sans-serif',
                'Tahoma, Geneva, sans-serif',
                'Times New Roman, Times, serif',
                'Trebuchet MS, Helvetica, sans-serif',
                'Verdana, Geneva, sans-serif'
            ],
            supportAllValues: true
        },
        fontSize: {
            options: [10, 12, 14, 'default', 18, 20, 22],
            supportAllValues: true
        },
        htmlSupport: {
            allow: [
                {
                    name: /.*/,
                    attributes: true,
                    classes: true,
                    styles: true
                }
            ]
        },
        htmlEmbed: {
            showPreviews: true
        },
        link: {
            decorators: {
                addTargetToExternalLinks: true,
                defaultProtocol: 'https://',
                toggleDownloadable: {
                    mode: 'manual',
                    label: 'Downloadable',
                    attributes: {
                        download: 'file'
                    }
                }
            }
        },
        mention: {
            feeds: [
                {
                    marker: '@',
                    feed: [
                        '@apple', '@bears', '@brownie', '@cake', '@cake', '@candy', '@canes', '@chocolate', '@cookie', '@cotton', '@cream',
                        '@cupcake', '@danish', '@donut', '@dragée', '@fruitcake', '@gingerbread', '@gummi', '@ice', '@jelly-o',
                        '@liquorice', '@macaroon', '@marzipan', '@oat', '@pie', '@plum', '@pudding', '@sesame', '@snaps', '@soufflé',
                        '@sugar', '@sweet', '@topping', '@wafer'
                    ],
                    minimumCharacters: 1
                }
            ]
        },
        removePlugins: [
            'ExportPdf',
            'ExportWord',
            'CKBox',
            'CKFinder',
            'EasyImage',
            'RealTimeCollaborativeComments',
            'RealTimeCollaborativeTrackChanges',
            'RealTimeCollaborativeRevisionHistory',
            'PresenceList',
            'Comments',
            'TrackChanges',
            'TrackChangesData',
            'RevisionHistory',
            'Pagination',
            'WProofreader',
            'MathType',
            'SlashCommand',
            'Template',
            'DocumentOutline',
            'FormatPainter',
            'TableOfContents',
            'PasteFromOfficeEnhanced'
        ]
    }).then(ckeditor => {
        console.log(Array.from(ckeditor.ui.componentFactory.names()));
        ckeditor1 = ckeditor
        ckeditor.plugins.get('FileRepository').createUploadAdapter = (loader) => {
            return {
                upload: () => {
                    return loader.file.then(file => {
                        return new Promise((resolve, reject) => {
                            const formData = new FormData();
                            formData.append('file', file);
                            formData.append('csrf', $("input[name='csrf']").val())
                            fetch(url + 'spaces/imageupload', {
                                method: 'POST',
                                body: formData
                            })
                                .then(response => response.json())
                                .then(data => {
                                    if (data.success) {
                                        resolve({
                                            default: data.url // URL to the uploaded image
                                        });
                                    } else {
                                        reject(data.error);
                                    }
                                })
                                .catch(error => {
                                    reject(error);
                                });
                        });
                    });
                }
            };
        };
        ckeditor.model.document.on('change', (event, data) => {
            KeyupContent()
        });
    })
        .catch(error => {
            console.error(error);
        });

}


/** Page String */
function AddPageString(name, newid, languagedata, id, orderindex, creupd) {

    var creorupd = creupd == true ? "viewpage" : "updpageview"

    var pagestr =
        `<div class="sort-index page${id}${newid}" data-id=${id} data-newid=${newid} data-collapcount=${collapsecount} data-orderindex=${orderindex}>
            <div class="spaceAccord accordionOption">
                <a class="spaceAccord-btn accordion-button ${creorupd}" type="button" data-bs-toggle="collapse" data-bs-target="#accord${collapsecount}" aria-expanded="true" aria-controls="collapse${collapsecount}">
                    <div style="display:flex">
                    <p class="para">${name}</p>
                    <img src="/public/img/accordion-vector.svg" style="display:none;"/>
                    </div>
                    <div class="card-option field-delete">
                        <div class="btn-group language-group space-option-grp">
                            <button type="button" class="dropdown-tog" data-bs-toggle="dropdown" aria-expanded="false">
                                <img src="/public/img/card-option.svg" alt="">               
                            </button>
                            <ul class="dropdown-menu dropdown-menu-end" data-popper-placement="bottom-end" style=" position: absolute; inset: 0px 0px auto auto; margin: 0px; transform: translate(0px, 22px);">
                                <li>
                                    <button class="dropdown-item insertsubpage" type="button">
                                     ` + languagedata.Spaces.insertsubpages + `
                                    </button>
                                </li>
                                <li>
                                    <button class="dropdown-item pageabove" type="button">
                                        ` + languagedata.Spaces.addpageabove + `
                                    </button>
                                </li>
    
                                <li>
                                    <button class="dropdown-item pagebelow" type="button">
                                        `+ languagedata.Spaces.addpagebelow + `
                                    </button>
                                </li>
                                <li>
                                    <button class="dropdown-item pageclone" type="button">
                                        ` + languagedata.Spaces.clone + `
                                    </button>
                                </li>
                                <li>
                                    <button class="dropdown-item movepage" type="button" data-bs-toggle="modal" data-bs-target="#moveModal">
                                        ` + languagedata.Spaces.movepage + `
                                    </button>
                                </li>
                                <li>
                                    <button class="dropdown-item deletepage" type="button" data-bs-toggle="modal" data-bs-target="#centerModal">
                                        ` + languagedata.delete + `
                                    </button>
                                </li>
                            </ul>
                        </div>
                    </div>
                </a> 
            </div>
        </div>`

    collapsecount++;

    return pagestr
}

function AddGroupString(dbid, newid, languagedata, grpname, orderindex) {

    var GroupName = grpname == "" ? "Group" + GroupId : grpname

    var grpstr = `
        <div class="groupsdiv${dbid} grpss" data-id=${dbid} data-newid=${newid} id=grp${dbid}${newid} data-orderindex=${orderindex}>
            <div class="newAccord-group">
                <textarea class="para groupnamee" style="resize:none;width:100%;overflow:hidden">${GroupName}</textarea>
                <div class="card-option field-delete">
                    <div class="btn-group language-group space-option-grp">
                        <button type="button" class="dropdown-tog" data-bs-toggle="dropdown" aria-expanded="false">
                            <img src="/public/img/card-option.svg" alt="">
                        </button>

                        <ul class="dropdown-menu dropdown-menu-end" data-popper-placement="bottom-end" style="position: absolute; inset: 0px 0px auto auto;margin: 0px;transform: translate(0px, 22px);">
                                <li>
                                    <button class="dropdown-item insertpage" type="button">
                                           `+ languagedata.Spaces.insertpage + `
                                    </button>
                                </li>
                                <li>
                                    <button class="dropdown-item grouppageabove" type="button">
                                    ` + languagedata?.Spaces?.addpageabove + `
                                    </button>
                                </li>

                                <li>
                                    <button class="dropdown-item grouppagebelow" type="button">
                                    ` + languagedata.Spaces.addpagebelow + `
                                    </button>
                                </li>

                                <li>
                                    <button class="dropdown-item groupdelete" type="button" data-bs-toggle="modal" data-bs-target="#centerModal">
                                    ` + languagedata.delete + `
                                    </button>
                                </li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    `

    return grpstr
}

function AddSubString(colcount, subpagename, spid, newspid) {

    var name = subpagename == "" ? `New Page ${SubPageId}` : subpagename

    var substr =
        `
    <div id="accord${colcount}" class="accordion-collapse card-option collapse show spaceAccord-content subpageoption subpageview subpag${spid}${newspid}" aria-labelledby="headingOne" data-bs-parent="#accordionExample" data-id=${spid} data-newid=${newspid}>
        <p class="para high">${name}</p>
        <div class="btn-group language-group space-option-grp">
            <button type="button" class="dropdown-tog" data-bs-toggle="dropdown" aria-expanded="true">
                <img src="/public/img/card-option.svg" alt="">               
            </button>
            <ul class="dropdown-menu dropdown-menu-end" data-popper-placement="bottom-end" style="position: absolute; inset: 0px 0px auto auto; margin: 0px; transform: translate(0px, 16.6667px);">
                <li>
                    <button class="dropdown-item subpageabove" type="button">
                        Add Page Above
                    </button>
                </li>
                <li>
                    <button class="dropdown-item subpagebelow" type="button">
                        Add Page Below
                    </button>
                </li>
                <li>
                    <button class="dropdown-item subpageclone" type="button">
                        Clone
                    </button>
                </li>
                <li>
                    <button class="dropdown-item subdeletepage" type="button" data-bs-toggle="modal" data-bs-target="#centerModal">
                        Delete
                    </button>
                </li>
            </ul>
        </div>
    </div>
    `
    return substr

}

/** View Page content */
$(document).on('click', '.viewpage', function () {

    var id = $(this).parents('.sort-index').attr('data-id');

    var newid = $(this).parents('.sort-index').attr('data-newid');

    $('.subpageview').removeClass('subpageselected');

    $('.viewpage').removeClass('pageselected');

    $('.updsubpageview').removeClass('subpageselected');

    $('.updpageview').removeClass('pageselected');

    $(this).addClass('pageselected');

    const pageindex = newpagesarr.findIndex(obj => {

        return obj.PageId == id && obj.NewPageId == newid;

    })

    if (pageindex >= 0) {

        $('#pagetitle').val(newpagesarr[pageindex].Name);

        ckeditor1.setData(newpagesarr[pageindex].Content);

        // $('#readtimetag').val(newpagesarr[pageindex].ReadTime)

        $('#updatedate').parents('.space-head-right').hide()

        // if(newpagesarr[pageindex].ReadTime>1){

        //     $('.minclass').text('mins')

        // }else{
    
        //     $('.minclass').text('min')
    
        // }

    }

    window.location.replace("#" + $('#pagetitle').val().replaceAll(" ", "_"));

    $('.space-section-rgt-body').html("");
})


/** Update Page content */
$(document).on('click', '.updpageview', function () {

    console.log(newpagesarr,"arrat");

    var id = $(this).parents('.sort-index').attr('data-id');

    var newid = $(this).parents('.sort-index').attr('data-newid');

    $('.updsubpageview').removeClass('subpageselected');

    $('.updpageview').removeClass('pageselected');

    $('.subpageview').removeClass('subpageselected');

    $('.viewpage').removeClass('pageselected');

    $(this).addClass('pageselected');

    const upageindex = updPagesarr.findIndex(obj => {

        return obj.PageId == id && obj.NewPageId == newid;

    })

    if (upageindex >= 0) {

        $('#pagetitle').val(updPagesarr[upageindex].Name)

        ckeditor1.setData(updPagesarr[upageindex].Content);

        $('#createdusername').text(updPagesarr[upageindex].Username);

        // $('#readtimetag').val(updPagesarr[pageindex].ReadTime)

        // if(updPagesarr[pageindex].ReadTime>1){

        //     $('.minclass').text('mins')

        // }else{
    
        //     $('.minclass').text('min')
    
        // }

        if ($(".page" + id + newid).children('div').children('a').hasClass('updpageview')) {

            console.log("date----", updPagesarr[upageindex]);

            var date = get_time_diff(updPagesarr[upageindex].LastUpdate);

            console.log(date, "---------");

            $('#updatedate').parents('.space-head-right').show()

            if (date == 0) {

                $('#updatedate').text(date + " minute ago")

            } else if (date < 60) {

                $('#updatedate').text(date + " minutes ago")

            } else {

                $('#updatedate').text(Layout(updPagesarr[upageindex].LastUpdate))

            }
        }


    } else {

        const pageindex = tempPagesarr.findIndex(obj => {

            return obj.PageId == id && obj.NewPageId == newid;

        })

        if (pageindex >= 0) {

            $('#pagetitle').val(tempPagesarr[pageindex].Name)

            ckeditor1.setData(tempPagesarr[pageindex].Content);

            $('#createdusername').text(tempPagesarr[pageindex].Username);

            // console.log(tempPagesarr[pageindex].ReadTime,"check")
            // $('#readtimetag').val(tempPagesarr[pageindex].ReadTime)

            // if(tempPagesarr[pageindex].ReadTime>1){

            //     $('.minclass').text('mins')
    
            // }else{
        
            //     $('.minclass').text('min')
        
            // }
    
            

            // $("#updatedate").text(tempPagesarr[pageindex].LastUpdate);

            if ($(".page" + id + newid).children('div').children('a').hasClass('updpageview')) {

                var date = get_time_diff(tempPagesarr[pageindex].LastUpdate);

                $('#updatedate').parents('.space-head-right').show()

                if (date == 0) {

                    $('#updatedate').text(date + " minute ago")

                } else if (date < 60) {

                    $('#updatedate').text(date + " minutes ago")

                } else {

                    $('#updatedate').text(Layout(tempPagesarr[pageindex].LastUpdate))

                }
            }




        }
    }


    const pageindex = tempPagesarr.findIndex(obj => {

        return obj.PageId == id && obj.NewPageId == newid;

    })

    if (tempPagesarr[pageindex].Log != null && tempPagesarr[pageindex].Log.length > 0) {

        var div = ""

        for (let x of tempPagesarr[pageindex].Log) {

            div += `
            <div class="space-rgt-bdy-row">       
                <div class="space-rgt-bdy-col">
                    <h3 class="para-light">${x.Username}</h3>
                    <p class="para-extralight">${Layout(x.Date)}</p>
                </div>
            <div class="space-rgt-bdy-col">`
            if (x.Status == "publish") {

                div += `<p class="para-extralight  status-published ">Published</p>`

            } else if (x.Status == "draft") {

                div += `<p class="para-extralight  status-draft" style ="color: #FC770F">Draft</p>`

            } else {
                div += `<p class="para-extralight  status-updated">${x.Status}</p>`

            }
            div += `</div>
        </div>`

        }
        $('.space-section-rgt-body').html(div);

    } else {
        $('.space-section-rgt-body').html("");
    }

    window.location.replace("#" + $('#pagetitle').val().replaceAll(" ", "_"));
})


/** View SubPage content */
$(document).on('click', '.subpageview', function () {

    var id = $(this).attr('data-id');

    var newid = $(this).attr('data-newid');

    $('.pageselected').removeClass('pageselected');

    $('.subpageview').removeClass('subpageselected');

    $(this).addClass('subpageselected');

    var page = FindSubPageName(id, newid)

    $('#pagetitle').val(page.Name)

    // console.log(page.ReadTime,"check1")

    // $('#readtimetag').val(page.ReadTime)

    // if(page.ReadTime > 1){

    //     $('.minclass').text('mins')
    
    // }else{

    //     $('.minclass').text('min')

    // }

    ckeditor1.setData(page.Content);

    window.location.replace("#" + $('#pagetitle').val().replaceAll(" ", "_"));


})

/** Keyup content save to object */
function KeyupContent() {

    console.log("content keyup");

    var data = ckeditor1.getData();

    //     var content = decodeHTML(data).replace(/<[^>]*>/g, '').replace(/&(?:nbsp|amp|quot|#39);/g, '');

    //    console.log(content,"newdata")

    var pageid = $('.pageselected').parents('.sort-index').attr('data-id');

    var newpageid = $('.pageselected').parents('.sort-index').attr('data-newid');

    var subpageid = $('.subpageselected').attr('data-id');

    var newsubpageid = $('.subpageselected').attr('data-newid');

    // const words = content.trim().split(/\s+/).length;

    // const time = Math.ceil(words / wpm);

    // console.log(time,words);

    if ($('.pageselected').hasClass('updpageview')) {

        if (pageid != 0 || newpageid != 0) {

            const updatepageindex = updPagesarr.findIndex(obj => {

                return obj.PageId == pageid && obj.NewPageId == newpageid;

            })

            if (updatepageindex >= 0) {

                if (updPagesarr[updatepageindex].Content != data || updPagesarr[updatepageindex].Name != $('#pagetitle').val()) {

                    $('.page' + pageid + newpageid).children('.accordionOption').children('a').addClass('pageupdatedd');

                }

                updPagesarr[updatepageindex].Content = data

            } else {


                const pageindex = tempPagesarr.findIndex(obj => {

                    return obj.PageId == pageid && obj.NewPageId == newpageid;

                })


                if (tempPagesarr[pageindex].Content != data || tempPagesarr[pageindex].Name != $('#pagetitle').val()) {

                    updPagesarr.push(tempPagesarr[pageindex]);

                    console.log(updPagesarr, "arry")

                    const updatepageindex1 = updPagesarr.findIndex(obj => {

                        return obj.PageId == pageid && obj.NewPageId == newpageid;

                    })

                    console.log(updatepageindex1, "")
                    if (updatepageindex1 >= 0) {

                        // if (updPagesarr[updatepageindex].Content != data || updPagesarr[updatepageindex].Name != $('#pagetitle').val()) {

                        $('.page' + pageid + newpageid).children('.accordionOption').children('a').addClass('pageupdatedd');

                        // }

                        updPagesarr[updatepageindex1].Content = data

                    }

                }


            }

        }

    } else {

        if (pageid != 0 || newpageid != 0) {

            const pageindex = newpagesarr.findIndex(obj => {

                return obj.PageId == pageid && obj.NewPageId == newpageid;

            })

            if (pageindex >= 0) {

                newpagesarr[pageindex].Content = data

            }

        }

        if (subpageid != 0 || newsubpageid != 0) {

            const pageindex = newSubpagearr.findIndex(obj => {

                return obj.SubPageId == subpageid && obj.NewSubPageId == newsubpageid;

            })

            if (pageindex >= 0) {

                newSubpagearr[pageindex].Content = data

            }

        }

    }

    if (subpageid != 0 || newsubpageid != 0) {

        const updatesubpageindex = updSubpagearr.findIndex(obj => {

            return obj.SubPageId == subpageid && obj.NewSubPageId == newsubpageid;

        })


        if (updatesubpageindex >= 0) {

            if (updSubpagearr[updatesubpageindex].Content != data || updSubpagearr[updatesubpageindex].Name != $('#pagetitle').val()) {

                $('.subpag' + subpageid + newsubpageid).addClass('pageupdatedd');

            }
            updSubpagearr[updatesubpageindex].Content = data

        } else {

            const pageindex = tempSubpagearr.findIndex(obj => {

                return obj.SubPageId == subpageid && obj.NewSubPageId == newsubpageid;

            })

            if (pageindex >= 0) {

                // if (tempSubpagearr[pageindex].Content != data || tempSubpagearr[pageindex].Name != $('#pagetitle').val()) {

                updSubpagearr.push(tempSubpagearr[pageindex]);

                // }
            }
        }

    }
}

/**Page above  */
$(document).on('click', '.pageabove', function () {

    var pagename = "New Page " + PageId

    var AddPage = AddPageString(pagename, PageId, languagedata, 0, OrderIndex, true);

    $(AddPage).insertBefore($(this).parents(".sort-index"));

    var pageobj = PageObjectCreate();

    pageobj.OrderIndex = OrderIndex

    pageobj.Name = pagename

    pageobj.PageGroupId = parseInt($(this).parents('.grpss').attr('data-id'))

    pageobj.NewGroupId = parseInt($(this).parents('.grpss').attr('data-newid'))

    console.log(pageobj);

    newpagesarr.push(pageobj)

    PageId++;

    OrderIndex++;
})
$(document).on('click', '.subpageabove', function () {

    console.log("orderindex", OrderIndex)

    var pagename = "New Page " + PageId

    var colcount = $(this).parents('.sort-index').attr('data-collapcount')

    var AddPage = AddSubString(colcount, "", 0, SubPageId);

    $(AddPage).insertBefore($(this).parents(".subpageview"));

    var pageobj = SubPageObjectCreate();

    pageobj.OrderIndex = OrderIndex

    pageobj.Name = pagename

    pageobj.PageGroupId = parseInt($(this).parents('.grpss').attr('data-id'))

    pageobj.NewGroupId = parseInt($(this).parents('.grpss').attr('data-newid'))

    pageobj.ParentId = parseInt($(this).parents('.sort-index').attr('data-id'))

    pageobj.NewParentId = parseInt($(this).parents('.sort-index').attr('data-newid'))

    newSubpagearr.push(pageobj)

    SubPageId++;

    OrderIndex++;
})
$(document).on('click', '.subpagebelow', function () {

    var pagename = "New Page " + PageId

    var colcount = $(this).parents('.sort-index').attr('data-collapcount')

    var AddPage = AddSubString(colcount, "", 0, SubPageId);

    $(AddPage).insertAfter($(this).parents(".subpageview"));

    var pageobj = SubPageObjectCreate();

    pageobj.OrderIndex = OrderIndex

    pageobj.Name = pagename

    pageobj.PageGroupId = parseInt($(this).parents('.grpss').attr('data-id'))

    pageobj.NewGroupId = parseInt($(this).parents('.grpss').attr('data-newid'))

    pageobj.ParentId = parseInt($(this).parents('.sort-index').attr('data-id'))

    pageobj.NewParentId = parseInt($(this).parents('.sort-index').attr('data-newid'))

    console.log(pageobj);

    newSubpagearr.push(pageobj)

    SubPageId++;

    OrderIndex++;
})

/**Group Page above  */
$(document).on('click', '.grouppageabove', function () {

    var pagename = "New Page " + PageId

    var AddPage = AddPageString(pagename, PageId, languagedata, 0, OrderIndex, true);

    $(AddPage).insertBefore($(this).parents(".grpss"));

    var pageobj = PageObjectCreate();

    pageobj.OrderIndex = OrderIndex

    pageobj.Name = pagename

    newpagesarr.push(pageobj)

    PageId++;

    OrderIndex++;
})

/**Page below  */
$(document).on('click', '.pagebelow', function () {

    var pagename = "New Page " + PageId

    var AddPage = AddPageString(pagename, PageId, languagedata, 0, OrderIndex, true);

    $(AddPage).insertAfter($(this).parents(".sort-index"));

    var pageobj = PageObjectCreate();

    pageobj.Name = pagename

    pageobj.OrderIndex = OrderIndex,

        pageobj.PageGroupId = parseInt($(this).parents('.grpss').attr('data-id')),

        pageobj.NewGroupId = parseInt($(this).parents('.grpss').attr('data-newid'))

    newpagesarr.push(pageobj)

    PageId++;

    OrderIndex++;
})


/**Group Page below  */
$(document).on('click', '.grouppagebelow', function () {

    var pagename = "New Page " + PageId

    var AddPage = AddPageString(pagename, PageId, languagedata, 0, OrderIndex, true);

    $(AddPage).insertAfter($(this).parents(".grpss"));

    var pageobj = PageObjectCreate();

    pageobj.Name = pagename

    pageobj.OrderIndex = OrderIndex,

        newpagesarr.push(pageobj)

    PageId++;

    OrderIndex++;
})

/** Move Page */
$(document).on('click', '.movepage', function () {

    $('#group-error').hide();

    $('.choosegroup').text('Choose the Page Group')

    var gid = $(this).parents('.grpss').attr('data-id')

    var newgid = $(this).parents('.grpss').attr('data-newid')

    var pageid = $(this).parents('.sort-index').attr('data-id')

    var newpageid = $(this).parents('.sort-index').attr('data-newid')

    var grps = ""

    var TotalAvaliableGroups = []

    TotalAvaliableGroups = TotalAvaliableGroups.concat(tempGrouparr, newGrouparr, updGrouparr)

    const unique = TotalAvaliableGroups.filter((obj, index) => {
        return index === TotalAvaliableGroups.findIndex(o => obj.GroupId === o.GroupId && obj.NewGroupId === o.NewGroupId);
    });

    for (let x of unique) {

        grps += `<button type="button"  class="dropdown-item groupclick" id="mgrp${x.GroupId}${x.NewGroupId}" data-id=${x.GroupId} data-newid=${x.NewGroupId}>${x.Name}</button>`
    }

    console.log(gid, newgid);

    $('.movegroupdiv').html(grps);


    if (gid != undefined && newgid != undefined) {

        $('#mgrp' + gid.toString() + newgid.toString()).trigger('click');
    }

    $('#pagedatass').attr('data-id', pageid).attr('data-newid', newpageid);

    $('#pgnam').val($(this).parents('.card-option').siblings('div').children('p').text())
})

$(document).on('click', '.groupclick', function () {

    var gid = $(this).attr('data-id');

    var newgid = $(this).attr('data-newid');

    $('#pagedatass').attr('data-groupid', gid).attr('data-newgroupid', newgid);

    $('#group-error').hide();


})

/** Clone Page */
$(document).on('click', '.pageclone', function () {

    var id = $(this).parents('.sort-index').attr('data-id');

    var newid = $(this).parents('.sort-index').attr('data-newid');

    PageClone(id, newid)



})

/** Group clone */
$(document).on('click', '.groupclone', function () {

    var id = $(this).parents('.grpss').attr('data-id');

    var newid = $(this).parents('.grpss').attr('data-newid');

    const grpindex = newGrouparr.findIndex(obj => {

        return obj.GroupId == id && obj.NewGroupId == newid;

    })

    if (grpindex >= 0) {

        var grpstring = AddGroupString(0, GroupId, languagedata, '', OrderIndex)

        $(grpstring).insertAfter($('#grp' + id + newid))

    }

})

/** Insert Page into group */
$(document).on('click', '.insertpage', function () {

    var id = $(this).parents('.grpss').attr('data-id');

    var newid = $(this).parents('.grpss').attr('data-newid');

    var pagename = "New Page " + PageId

    var pagedata = AddPageString(pagename, PageId, languagedata, 0, OrderIndex, true)

    $('#grp' + id + newid).append(pagedata)

    var pageobj = PageObjectCreate();

    pageobj.PageGroupId = parseInt(id)

    pageobj.NewGroupId = parseInt(newid)

    pageobj.OrderIndex = OrderIndex

    pageobj.NewGroupId = parseInt(newid)

    pageobj.Name = pagename

    newpagesarr.push(pageobj)

    console.log(newpagesarr);

    PageId++;

    OrderIndex++;

})

/** keyup name change */
$(document).on('keyup', "#pagetitle", function () {

    var id = $('.pageselected').parents('.sort-index').attr('data-id');

    var newid = $('.pageselected').parents('.sort-index').attr('data-newid');

    var subid = $('.subpageselected').attr('data-id');

    var newsubid = $('.subpageselected').attr('data-newid');

    var value = $(this).val();

    if ($('.pageselected').hasClass('updpageview')) {

        if (id != 0 || newid != 0) {

            $('.pageselected').children('div').children('p').text(value);

            SetDataPage(id, newid, "", value, "")

            $('.pageselected').addClass("pageupdatedd")

            // $("<span>(DRAFT)</span>").insertAfter($('.pageselected').children('div').children('p'));

        }

    } else {

        if (id != 0 || newid != 0) {

            $('.pageselected').children('div').children('p').text(value);

            const pageindex = newpagesarr.findIndex(obj => {

                return obj.PageId == id && obj.NewPageId == newid;

            })

            if (pageindex >= 0) {

                newpagesarr[pageindex].Name = value

            }
        }

        if (subid != 0 || newsubid != 0) {

            $('.subpageselected').children('p').text(value);

            SetSubpageData(subid, newsubid, " ", value)

        }

    }

    if (subid != 0 || newsubid != 0) {

        $('.subpageselected').children('p').text(value);

        $('.subpageselected').addClass("pageupdatedd")

        SetSubpageData(subid, newsubid, "", value)

    }

})

/*Delete page Confirmation */
$(document).on('click', '.deletepage', function () {

    var pageid = $(this).parents(".sort-index").attr("data-id")

    var newpgid = $(this).parents(".sort-index").attr("data-newid")

    var page = FindPageName(pageid, newpgid)

    $("#delid").attr("data-pageid", pageid)

    $('#delid').attr('data-newpageid', newpgid)

    $('.deltitle').text(languagedata.Spaces.delpagetitle)

    $('.deldesc').text(languagedata.Spaces.delpagecontent)

    $('#content').text(languagedata.Spaces.delpagecontent)

    $("#delid").parent('#delete').attr('href', 'javascript:void(0);')

    $('.delname').text(page.Name)


})


/*Delete page Confirmation */
$(document).on('click', '.subdeletepage', function () {

    var pageid = $(this).parents(".subpageview").attr("data-id")

    var newpgid = $(this).parents(".subpageview").attr("data-newid")

    var Sub = FindSubPageName(pageid, newpgid)

    console.log("sub----", Sub);

    $("#delid").attr("data-subpageid", pageid)

    $('#delid').attr('data-newsubpageid', newpgid)

    $('.deltitle').text(languagedata.Spaces.delpagetitle)

    $('.deldesc').text(languagedata.Spaces.delpagecontent)

    $('.delname').text(Sub.Name)

    $('#content').text(languagedata.Spaces.delpagecontent)

    $("#delid").parent('#delete').attr('href', 'javascript:void(0);')

})


/*Delete page Confirmation */
$(document).on('click', '.groupdelete', function () {

    var gid = $(this).parents(".grpss").attr("data-id")

    var newgid = $(this).parents(".grpss").attr("data-newid")

    var name = FindGroupName(gid, newgid)

    $("#delid").attr("data-groupid", gid)

    $('#delid').attr('data-newgroupid', newgid)

    $('.deltitle').text("Delete Group?")

    $('.deldesc').text("Are you sure you want to delete this Group?")

    $('.delname').text(name)

    $("#delid").parent('#delete').attr('href', 'javascript:void(0);')

})


/**Add Page */
$(document).on('click', '#addpage', function () {

    var pagename = "New Page " + PageId

    var pagedata = AddPageString(pagename, PageId, languagedata, 0, OrderIndex, true)

    $('.newpages').append(pagedata)

    var pageobj = PageObjectCreate();

    pageobj.Name = pagename

    pageobj.OrderIndex = OrderIndex

    newpagesarr.push(pageobj)

    PageId++;

    OrderIndex++;

})

$(document).on('click', '#addgroup', function () {

    var grpstring = AddGroupString(0, GroupId, languagedata, '', OrderIndex)

    $('.newpages').append(grpstring)

    var obj = GroupCreationObject()

    obj.OrderIndex = OrderIndex

    obj.Name = "Group" + GroupId

    newGrouparr.push(obj)

    var pagename = "New Page " + PageId

    var pagedata = AddPageString(pagename, PageId, languagedata, 0, OrderIndex, true)

    $('#grp' + 0 + GroupId.toString()).append(pagedata)

    var pageobj = PageObjectCreate()

    pageobj.OrderIndex = OrderIndex,

        pageobj.NewGroupId = GroupId

    pageobj.Name = pagename

    newpagesarr.push(pageobj)

    PageId++;

    OrderIndex++;

    GroupId++;

})

/** Page Creation Object */
function PageObjectCreate() {

    var obj = {}

    obj.PageId = 0

    obj.NewPageId = PageId

    obj.Content = ""

    obj.PageGroupId = 0

    obj.NewGroupId = 0

    obj.ParentId = 0

    obj.ReadTime = 0

    return obj

}


/** SubPage Creation Object */
function SubPageObjectCreate() {

    var obj = {}

    obj.SubPageId = 0

    obj.NewSubPageId = SubPageId

    obj.Content = ""

    obj.PageGroupId = 0

    obj.ReadTime = 0

    // obj.NewGroupId = parseInt(newgrpid == undefined ? 0 : newgrpid)

    obj.OrderIndex = OrderIndex

    // obj.ParentId= parseInt(parentid)

    return obj

}


/** Group Creation Object */
function GroupCreationObject() {

    var obj = {}

    obj.GroupId = 0

    obj.NewGroupId = GroupId

    //  obj.Name = name.toUpperCase()

    return obj

}



// Savebtn
$(document).on('click', '#savebtn', function () {

    var spid = $("#spid").val();
    console.log("space", spid)

    var flg = ContentValidation()

    console.log(flg);

    if (!flg) {

        $.ajax({
            url: "/spaces/pagecreate",
            type: "POST",
            dataType: "json",
            data: {
                "spaceid": spid,
                "creategroups": JSON.stringify({ newGrouparr }),
                "createpages": JSON.stringify({ newpagesarr }),
                "createsubpage": JSON.stringify({ newSubpagearr }),
                "updategroups": JSON.stringify({ updGrouparr }),
                "updatepages": JSON.stringify({ updPagesarr }),
                "updatesubpage": JSON.stringify({ updSubpagearr }),
                "deletepage": JSON.stringify({ deletePagearr }),
                "deletegroup": JSON.stringify({ deletegrouparr }),
                "deletesub": JSON.stringify({ deletesubarr }),
                "save": "save",
                csrf: $("input[name='csrf']").val()
            },
            success: function (result) {
                if (result) {

                    setCookie("get-toast", "pagesaved")
                    setCookie('Alert-msg', 'success', 1)

                } else {

                    setCookie("Alert-msg", "Internal Server Error")
                }

                window.location.reload(true);

            }
        })
    }

})


$(document).on('click', '#publishbtn', function () {

    var spid = $("#spid").val();
    console.log("space", spid)


    var flg = ContentValidation()

    console.log(flg);

    if (!flg) {

        $.ajax({
            url: "/spaces/pagecreate",
            type: "POST",
            dataType: "json",
            data: {
                "spaceid": spid,
                "creategroups": JSON.stringify({ newGrouparr }),
                "createpages": JSON.stringify({ newpagesarr }),
                "createsubpage": JSON.stringify({ newSubpagearr }),
                "updategroups": JSON.stringify({ updGrouparr }),
                "updatepages": JSON.stringify({ updPagesarr }),
                "updatesubpage": JSON.stringify({ updSubpagearr }),
                "deletepage": JSON.stringify({ deletePagearr }),
                "deletegroup": JSON.stringify({ deletegrouparr }),
                "deletesub": JSON.stringify({ deletesubarr }),
                "save": "publish",
                csrf: $("input[name='csrf']").val()
            },
            success: function (result) {
                if (result) {

                    setCookie('get-toast', "pagePublished")
                    setCookie('Alert-msg', 'success', 1)

                } else {

                    setCookie("Alert-msg", "Internal Server Error")
                }

                window.location.reload(true);

            }
        })

    }
})
/*Insert Sub Page */
$(document).on('click', '.insertsubpage', function () {

    var colcount = $(this).parents('.sort-index').attr('data-collapcount')

    var Subpage = AddSubString(colcount, "", 0, SubPageId);

    $(this).parents('.spaceAccord').append(Subpage);

    var obj = SubPageObjectCreate();

    obj.Name = "New Page" + SubPageId

    newSubpagearr.push(obj)

    console.log(newSubpagearr);

    SubPageId++;

    obj.ParentId = parseInt($(this).parents('.sort-index').attr('data-id'));

    obj.NewParentId = parseInt($(this).parents('.sort-index').attr('data-newid'));

    $(this).parents('.spaceAccord-btn').children('div').children('p').siblings('img').show();

})

/** revision close */
$(document).on('click', '#revisionclose', function () {

    $(this).parents('.space-section-right').hide();
})

/**revision open */
$(document).on('click', '.revision-btn', function () {

    $('.space-section-right').show();

})


function PGList() {

    for (let x of overallarray) {

        /**this page */
        if (x['PgId'] != 0 && x['Pgroupid'] == 0) {

            var AddPage = AddPageString(x['Name'], 0, languagedata, x['PgId'], x['OrderIndex'], false);

            $('.newpages').append(AddPage)

            for (let j of tempSubpagearr) {

                if (j['ParentId'] == x['PgId']) {

                    var collap = $('.page' + j['ParentId'] + "0").attr('data-collapcount')

                    var AddSubPage = AddSubString(collap, j['Name'], j['SubPageId'], j['NewSubPageId'])

                    $('.page' + j['ParentId'] + "0").children('.accordionOption').append(AddSubPage)

                    $('.page' + j['ParentId'] + "0").children('.accordionOption').children('.spaceAccord-btn').children('div').children('p').siblings('img').show();

                }

            }

        }

        /**this Group */
        if (x['GroupId'] !== undefined && x['GroupId'] != 0 && x['PgId'] === undefined) {

            var AddGroup1 = AddGroupString(x['GroupId'], x['NewGroupId'], languagedata, x['Name'], x['OrderIndex'])

            $('.newpages').append(AddGroup1)

            for (let y of overallarray) {

                if ((x['GroupId'] == y['Pgroupid']) && y['GroupId'] === undefined) {

                    var AddPage = AddPageString(y['Name'], 0, languagedata, y['PgId'], y['OrderIndex'], false)

                    $('.groupsdiv' + x['GroupId']).append(AddPage)

                    for (let j of tempSubpagearr) {

                        if (j['ParentId'] == y['PgId'] && y['Pgroupid'] != 0) {

                            var collap = $('.page' + j['ParentId'] + "0").attr('data-collapcount')

                            var AddSubPage = AddSubString(collap, j['Name'], j['SubPageId'], j['NewSubPageId'])

                            $('.page' + j['ParentId'] + "0").children('.accordionOption').append(AddSubPage)

                            $('.page' + j['ParentId'] + "0").children('.accordionOption').children('.spaceAccord-btn').children('div').children('p').siblings('img').show();

                        }

                    }
                }

            }

        }
    }

    $('.groupnamee').each(function () {
        $(this).css('height', 'auto').css('height', this.scrollHeight + (this.offsetHeight - this.clientHeight));
    })

    DefaultOrderIndex()
}

$(document).on('keyup', ".groupnamee", function () {

    var name = $(this).val();

    var id = $(this).parents('.grpss').attr('data-id')

    var newid = $(this).parents('.grpss').attr('data-newid')

    SetGroupName(id, newid, name)
})

function PageSelect() {

    $('.spaceAccord-btn').each(function () {

        $(this).trigger('click');

        return false;
    })
}

$(document).on('click', "#pagecancel", function () {

    $('.deltitle').text("Exit Pages?");

    $('.deldesc').text("Are you sure you want to Exit?");

    $("#delid").parent('#delete').attr('href', '/spaces/');

    $('.delname').text("");

})




/** Delete Page & Group & Subpage */
$(document).on('click', '#delid', function () {

    var pageid = $(this).attr('data-pageid');

    var newpageid = $(this).attr('data-newpageid');

    var groupid = $(this).attr('data-groupid');

    var newgroupid = $(this).attr('data-newgroupid');

    var subpageid = $(this).attr('data-subpageid');

    var newsubpageid = $(this).attr('data-newsubpageid');

    /**delete page */
    if (pageid != undefined || newpageid != undefined) {

        DeletePageFromArray(pageid, newpageid)

    }

    /** delete group */
    if (groupid != undefined || newgroupid != undefined) {

        DeleteGroupFromArray(groupid, newgroupid)
    }

    /** delete subpage */
    if (subpageid != undefined) {

        DeleteSubFromArray(subpageid, newsubpageid)

    }

    $('#delcancel').trigger('click');

    console.log(deletePagearr);
    console.log(newpagesarr);

    PageSelect()
})


/**DeletGroup */
function DeleteGroupFromArray(groupid, newgroupid) {

    if (groupid != 0) {

        const tempgroupindex = tempGrouparr.findIndex(obj => {

            return obj.GroupId == groupid && obj.NewGroupId == newgroupid;

        })

        if (tempgroupindex >= 0) {

            deletegrouparr.push(tempGrouparr[tempgroupindex])

            const updgroupindex = updGrouparr.findIndex(obj => {

                return obj.GroupId == groupid && obj.NewGroupId == newgroupid;

            })

            if (updgroupindex >= 0) {

                updGrouparr.splice(updgroupindex, 1)

            } else {

                const newgroupindex = newGrouparr.findIndex(obj => {

                    return obj.GroupId == groupid && obj.NewGroupId == newgroupid;

                })

                if (newgroupindex >= 0) {

                    newGrouparr.splice(newgroupindex, 1)

                }

            }


            DeletePageFromGroupid(groupid, newgroupid)

        }

    } else {

        const groupindex = newGrouparr.findIndex(obj => {

            return obj.GroupId == groupid && obj.NewGroupId == newgroupid;

        })

        if (groupindex >= 0) {

            deletegrouparr.push(newGrouparr[groupindex])

            newGrouparr.splice(groupindex, 1)

            DeletePageFromGroupid(groupid, newgroupid)

        }

    }

    console.log(deletePagearr);
    console.log(deletegrouparr);
    console.log(deletesubarr)

    $("#grp" + groupid + newgroupid).remove();

}

/** Delete page by page */
function DeletePageFromGroupid(groupid, newgroupid) {

    console.log("enter");

    for (let [index, obj] of newpagesarr.entries()) {

        console.log(obj.PageGroupId, groupid, obj.NewGroupId, newgroupid);

        if (obj.PageGroupId == groupid && obj.NewGroupId == newgroupid) {

            deletePagearr.push(newpagesarr[index])

            newpagesarr.slice(index, 1)

            DeleteSubPageByParentId(obj.PageId, obj.NewPageId)

        }

    }

    for (let [index, obj] of updPagesarr.entries()) {

        if (obj.PageGroupId == groupid && obj.NewGroupId == newgroupid) {

            deletePagearr.push(updPagesarr[index])

            updPagesarr.splice(index, 1)

            DeleteSubPageByParentId(obj.PageId, obj.NewPageId)



        }
    }

    for (let [index, obj] of tempPagesarr.entries()) {

        if (obj.PageGroupId == groupid && obj.NewGroupId == newgroupid) {

            deletePagearr.push(tempPagesarr[index])

            DeleteSubPageByParentId(obj.PageId, obj.NewPageId)

        }
    }

}

/** Delete Page */
function DeletePageFromArray(pageid, newpageid) {

    for (let [index, obj] of newpagesarr.entries()) {

        console.log(obj.PageId, pageid, obj.NewPageId, newpageid);

        if (obj.PageId == pageid && obj.NewPageId == newpageid) {

            deletePagearr.push(newpagesarr[index])

            newpagesarr.splice(index, 1)

        }

    }

    for (let [index, obj] of updPagesarr.entries()) {

        if (obj.PageId == pageid && obj.NewPageId == newpageid) {

            deletePagearr.push(updPagesarr[index])

            updPagesarr.splice(index, 1)

        }
    }

    for (let [index, obj] of tempPagesarr.entries()) {

        if (obj.PageId == pageid && obj.NewPageId == newpageid) {

            deletePagearr.push(tempPagesarr[index])
        }
    }

    $('.page' + pageid + newpageid).remove();

}

/** Delete subpage by parentid */
function DeleteSubPageByParentId(parentid, newparentid) {

    for (let [index, obj] of newSubpagearr.entries()) {

        if (obj.ParentId == parentid && obj.NewParentId == newparentid) {

            deletesubarr.push(newSubpagearr[pageindex])

            newSubpagearr.splice(index, 1)

        }

    }

    for (let [index, obj] of updSubpagearr.entries()) {

        if (obj.ParentId == parentid && obj.NewParentId == newparentid) {

            deletesubarr.push(updSubpagearr[pageindex])

            updSubpagearr.splice(index, 1)

        }
    }

}

/** Delete Sub */
function DeleteSubFromArray(subpageid, newsubpageid) {

    for (let [index, obj] of newSubpagearr.entries()) {

        if (obj.SubPageId == subpageid && obj.NewSubPageId == newsubpageid) {

            deletesubarr.push(newSubpagearr[index])

            newSubpagearr.splice(index, 1)

        }

    }

    for (let [index, obj] of updSubpagearr.entries()) {

        if (obj.SubPageId == subpageid && obj.NewSubPageId == newsubpageid) {

            deletesubarr.push(updSubpagearr[index])

            updSubpagearr.splice(index, 1)

        }
    }

    for (let [index, obj] of tempSubpagearr.entries()) {

        if (obj.SubPageId == subpageid && obj.NewSubPageId == newsubpageid) {

            deletesubarr.push(tempSubpagearr[index])

        }
    }


    $('.subpag' + subpageid + newsubpageid).remove();

}

/**Find page name */
function FindPageName(pageid, newpageid) {

    for (let [index, obj] of newpagesarr.entries()) {

        if (obj.PageId == pageid && obj.NewPageId == newpageid) {

            return newpagesarr[index]

        }

    }

    for (let [index, obj] of updPagesarr.entries()) {

        if (obj.PageId == pageid && obj.NewPageId == newpageid) {

            return updPagesarr[index]

        }
    }

    for (let [index, obj] of tempPagesarr.entries()) {

        if (obj.PageId == pageid && obj.NewPageId == newpageid) {

            return tempPagesarr[index]

        }

    }

}


/**Find page name */
function FindSubPageName(subpageid, newsubpageid) {

    const pageindex = tempSubpagearr.findIndex(obj => {

        return obj.SubPageId == subpageid && obj.NewSubPageId == newsubpageid;

    })

    if (pageindex >= 0) {

        if (tempSubpagearr[pageindex].Log != null && tempSubpagearr[pageindex].Log.length > 0 && pageindex >= 0) {


            var div = ""

            for (let x of tempSubpagearr[pageindex].Log) {

                div += `
                <div class="space-rgt-bdy-row">       
                    <div class="space-rgt-bdy-col">
                        <h3 class="para-light">${x.Username}</h3>
                        <p class="para-extralight">${Layout(x.Date)}</p>
                    </div>
                <div class="space-rgt-bdy-col">`
                if (x.Status == "publish") {

                    div += `<p class="para-extralight  status-published ">Published</p>`

                } else if (x.Status == "draft") {

                    div += `<p class="para-extralight  status-draft" style ="color: #FC770F">Draft</p>`

                } else {
                    div += `<p class="para-extralight  status-updated">${x.Status}</p>`

                }
                div += `</div>
            </div>`

            }
            $('.space-section-rgt-body').html(div);

        } else {

            $('.space-section-rgt-body').html("");

        }
    }



    for (let [index, obj] of updSubpagearr.entries()) {

        if (obj.SubPageId == subpageid && obj.NewSubPageId == newsubpageid) {

            console.log("upd", updSubpagearr[index]);

            var date = get_time_diff(updSubpagearr[index].LastUpdate);

            $('#updatedate').parents('.space-head-right').show()

            if (date == 0) {

                $('#updatedate').text(date + " minute ago")

            } else if (date < 60) {

                $('#updatedate').text(date + " minutes ago")

            } else {

                $('#updatedate').text(Layout(updSubpagearr[index].LastUpdate))

            }

            return updSubpagearr[index]

        }
    }


    for (let [index, obj] of newSubpagearr.entries()) {

        if (obj.SubPageId == subpageid && obj.NewSubPageId == newsubpageid) {

            console.log("new", newSubpagearr[index]);

            return newSubpagearr[index]

        }

    }


    for (let [index, obj] of tempSubpagearr.entries()) {

        if (obj.SubPageId == subpageid && obj.NewSubPageId == newsubpageid) {


            var date = get_time_diff(tempSubpagearr[index].LastUpdate);

            $('#updatedate').parents('.space-head-right').show()

            if (date == 0) {

                $('#updatedate').text(date + " minute ago")

            } else if (date < 60) {

                $('#updatedate').text(date + " minutes ago")

            } else {


                $('#updatedate').text(Layout(tempSubpagearr[index].LastUpdate))

            }

            return tempSubpagearr[index]

        }

    }

}

/**Find page name */
function FindGroupName(groupid, newgroupid) {

    for (let [index, obj] of tempGrouparr.entries()) {

        if (obj.GroupId == groupid && obj.NewGroupId == newgroupid) {

            return tempGrouparr[index].Name

        }

    }

    for (let [index, obj] of newGrouparr.entries()) {

        if (obj.GroupId == groupid && obj.NewGroupId == newgroupid) {

            return newGrouparr[index].Name

        }

    }

    for (let [index, obj] of updGrouparr.entries()) {

        if (obj.GroupId == groupid && obj.NewGroupId == newgroupid) {

            return updGrouparr[index].Name

        }
    }

}

/** Set Data */
function SetDataPage(id, newid, name, value, groupid) {

    console.log(value);

    const updatepageindex = updPagesarr.findIndex(obj => {

        return obj.PageId == id && obj.NewPageId == newid;

    })

    if (updatepageindex >= 0) {


        const updatepageindex = updPagesarr.findIndex(obj => {

            return obj.PageId == id && obj.NewPageId == newid;

        })

        if (updatepageindex >= 0) {

            if (value != "") {

                updPagesarr[updatepageindex].Name = value

                console.log(updPagesarr[updatepageindex]);
            }

            if (groupid['id'] != "" && groupid['newid'] != "" && groupid['id'] != undefined && groupid['newid'] != undefined) {

                updPagesarr[updatepageindex].PageGroupId = parseInt(groupid['id'])

                updPagesarr[updatepageindex].NewGroupId = parseInt(groupid['newid'])

            }

        } else {

            const npageindex = newpagesarr.findIndex(obj => {

                return obj.PageId == id && obj.NewPageId == newid;

            })

            if (npageindex >= 0) {

                if (groupid['id'] != "" && groupid['newid'] != "") {

                    newpagesarr[npageindex].GroupId = groupid['id']

                    newpagesarr[npageindex].NewGroupId = groupid['newid']

                }

            }
        }

    } else {

        const pageindex = tempPagesarr.findIndex(obj => {

            return obj.PageId == id && obj.NewPageId == newid;

        })

        if (pageindex >= 0) {

            updPagesarr.push(tempPagesarr[pageindex]);

            console.log(tempPagesarr[pageindex]);


            const updatepageindex = updPagesarr.findIndex(obj => {

                return obj.PageId == id && obj.NewPageId == newid;

            })

            if (updatepageindex >= 0) {

                updPagesarr[updatepageindex].Name = value

            }

            if (updatepageindex >= 0 && groupid['id'] != "" && groupid['newid'] != "" && groupid['id'] != undefined && groupid['newid'] != undefined) {

                updPagesarr[updatepageindex].PageGroupId = parseInt(groupid['id'])

                updPagesarr[updatepageindex].NewGroupId = parseInt(groupid['newid'])
            }
        }

    }

    console.log(updPagesarr);

    console.log(newpagesarr);

}

/** Set SubPageData */
function SetSubpageData(id, newid, name, value) {

    const upageindex = updSubpagearr.findIndex(obj => {

        return obj.SubPageId == id && obj.NewSubPageId == newid;

    })

    if (upageindex < 0) {

        const pageindex = tempSubpagearr.findIndex(obj => {

            return obj.SubPageId == id && obj.NewSubPageId == newid;

        })

        if (pageindex >= 0) {

            updSubpagearr.push(tempSubpagearr[pageindex]);

        } else {

            const npageindex = newSubpagearr.findIndex(obj => {

                return obj.SubPageId == id && obj.NewSubPageId == newid;

            })

            if (npageindex >= 0) {

                newSubpagearr[npageindex].Name = value;

            }
        }


    } else {

        updSubpagearr[upageindex].Name = value

        if (value != "") {

            updSubpagearr[upageindex].Name = value

        }

    }

}

/** Set Group Name */
function SetGroupName(id, newid, name) {

    const upgroupindex = updGrouparr.findIndex(obj => {

        return obj.GroupId == id && obj.NewGroupId == newid;

    })

    if (upgroupindex < 0) {

        const groupindex = tempGrouparr.findIndex(obj => {

            return obj.GroupId == id && obj.NewGroupId == newid;

        })


        if (groupindex >= 0) {

            updGrouparr.push(tempGrouparr[groupindex])

            const pgroupindex = updGrouparr.findIndex(obj => {

                return obj.GroupId == id && obj.NewGroupId == newid;

            })

            if (pgroupindex >= 0) {

                updGrouparr[pgroupindex].Name = name
            }

        } else {

            const ngroupindex = newGrouparr.findIndex(obj => {

                return obj.GroupId == id && obj.NewGroupId == newid;

            })

            if (ngroupindex >= 0) {

                newGrouparr[ngroupindex].Name = name
            }


        }

    } else {

        updGrouparr[upgroupindex].Name = name

    }

}


/** Keyup content save to object */
$(document).on("DOMSubtreeModified", ".newpages", function () {

    DefaultOrderIndex()

})

/** order index */
function DefaultOrderIndex() {

    var defaultorder = 1

    $('.newpages').children('div').each(function () {

        if ($(this).hasClass('grpss')) {

            $(this).attr('data-orderindex', defaultorder)

            $(this).children('.sort-index').each(function () {

                $(this).attr('data-orderindex', defaultorder)

                updatepageorderindex($(this).attr('data-id'), $(this).attr('data-newid'), defaultorder)

                defaultorder++
            })

            updategrouporderindex($(this).attr('data-id'), $(this).attr('data-newid'), defaultorder)

        } else {

            $(this).attr('data-orderindex', defaultorder)

            updatepageorderindex($(this).attr('data-id'), $(this).attr('data-newid'), defaultorder)

        }


        defaultorder++
    })

}

/**update group orderindex */
function updategrouporderindex(groupid, newgroupid, index) {

    const groupindex = newGrouparr.findIndex(obj => {

        return obj.GroupId == groupid && obj.NewGroupId == newgroupid;

    })

    if (groupindex >= 0) {

        newGrouparr[groupindex].OrderIndex = index

    }

    const upgroupindex = updGrouparr.findIndex(obj => {

        return obj.GroupId == groupid && obj.NewGroupId == newgroupid;

    })

    if (upgroupindex >= 0) {

        updGrouparr[groupindex].OrderIndex = index

    }

}


/**update group orderindex */
function updatepageorderindex(id, newid, index) {

    const pageindex = newpagesarr.findIndex(obj => {

        return obj.PageId == id && obj.NewPageId == newid;

    })

    if (pageindex >= 0) {

        newpagesarr[pageindex].OrderIndex = index

    }

    const uppageindex = updPagesarr.findIndex(obj => {

        return obj.PageId == id && obj.NewPageId == newid;

    })

    if (uppageindex >= 0) {

        updPagesarr[uppageindex].OrderIndex = index

    }

}

/** MovePage */
$(document).on('click', '#movepagetogroup', function () {

    var pid = $("#pagedatass").attr('data-id')

    var newpid = $("#pagedatass").attr('data-newid')

    var gid = $("#pagedatass").attr('data-groupid')

    var newgroupid = $("#pagedatass").attr('data-newgroupid')

    if (gid != 0 || newgroupid != 0 && (gid != undefined && newgroupid != undefined)) {

        var clonedElement = $('.page' + pid + newpid).clone();

        $('.page' + pid + newpid).remove();

        $("#grp" + gid + newgroupid).append(clonedElement);

        $('#movecancel').trigger('click');

        var obj = {}

        obj.id = gid

        obj.newid = newgroupid

        SetDataPage(pid, newpid, "", $('#pgnam').val(), obj)

        $('#pgnam').val('');

    } else {

        $('#group-error').show();

    }

})

function get_time_diff(datetime) {

    var givenTime = new Date(datetime);

    // Current time
    var currentTime = new Date();

    // Calculate the difference in milliseconds
    var timeDifference = currentTime - givenTime;

    // Convert milliseconds to minutes
    var timeDifferenceInMinutes = timeDifference / (1000 * 60);

    return Math.trunc(timeDifferenceInMinutes)
}

function Layout(datetime) {

    var givenTime = new Date(datetime);

    var options = { year: 'numeric', month: 'short', day: '2-digit', hour: '2-digit', minute: '2-digit' };

    // Convert the given timestamp to the desired format
    var formattedDate = givenTime.toLocaleString('en-US', options);

    var finaldate = formattedDate.replaceAll(",", "")

    return finaldate

}

function PageClone(id, newid) {

    const updpageindex = updPagesarr.findIndex(obj => {

        return obj.PageId == id && obj.NewPageId == newid;

    })

    const cpageindex = newpagesarr.findIndex(obj => {

        return obj.PageId == id && obj.NewPageId == newid;

    })

    var relatedsub = FindSubPageArray(id, newid)

    if (updpageindex >= 0) {

        var AddPage = AddPageString(updPagesarr[updpageindex].Name, PageId, languagedata, 0, OrderIndex, true);

        $(AddPage).insertAfter($('.page' + id + newid));

        var pageobj = PageObjectCreate();

        pageobj.OrderIndex = OrderIndex

        pageobj.Name = updPagesarr[updpageindex].Name

        pageobj.Content = updPagesarr[updpageindex].Content

        pageobj.PageGroupId = updPagesarr[updpageindex].GroupId

        pageobj.NewGroupId = updPagesarr[updpageindex].NewGroupId

        newpagesarr.push(pageobj)

        for (let x of relatedsub) {

            console.log("xx--", x);

            var colcount = $('.page' + id + newid).attr('data-collapcount')

            var Subpage = AddSubString(colcount, x.Name, 0, SubPageId);

            $('.page0' + PageId).children('.accordionOption').append(Subpage);

            $('.page0' + PageId).children('.spaceAccord').children('a').children('div').children('img').show();

            var obj = SubPageObjectCreate();

            obj.Name = x.Name

            obj.Content = x.Content

            obj.NewParentId = PageId

            newSubpagearr.push(obj)

            SubPageId++;
        }


    } else if (cpageindex >= 0) {

        var AddPage = AddPageString(newpagesarr[cpageindex].Name, PageId, languagedata, 0, OrderIndex, true);

        $(AddPage).insertAfter($('.page' + id + newid));

        var pageobj = PageObjectCreate();

        pageobj.OrderIndex = OrderIndex

        pageobj.Name = newpagesarr[cpageindex].Name

        pageobj.Content = newpagesarr[cpageindex].Content

        pageobj.PageGroupId = newpagesarr[cpageindex].GroupId

        pageobj.NewGroupId = newpagesarr[cpageindex].NewGroupId

        newpagesarr.push(pageobj)

        for (let x of relatedsub) {

            var colcount = $('.page0' + PageId).attr('data-collapcount')

            var Subpage = AddSubString(colcount, x.Name, 0, SubPageId);

            $('.page0' + PageId).children('.accordionOption').append(Subpage);

            $('.page0' + PageId).children('.spaceAccord').children('a').children('div').children('img').show();

            var obj = SubPageObjectCreate();

            obj.Name = x.Name

            obj.Content = x.Content

            obj.NewParentId = PageId

            newSubpagearr.push(obj)

            SubPageId++;
        }


    } else {

        const pageindex = tempPagesarr.findIndex(obj => {

            return obj.PageId == id && obj.NewPageId == newid;

        })

        if (pageindex >= 0) {

            var AddPage = AddPageString(tempPagesarr[pageindex].Name, PageId, languagedata, 0, OrderIndex, true);

            $(AddPage).insertAfter($('.page' + id + newid));

            var pageobj = PageObjectCreate();

            pageobj.OrderIndex = OrderIndex

            pageobj.Name = tempPagesarr[pageindex].Name

            pageobj.Content = tempPagesarr[pageindex].Content

            pageobj.PageGroupId = tempPagesarr[pageindex].GroupId

            pageobj.NewGroupId = tempPagesarr[pageindex].NewGroupId

            newpagesarr.push(pageobj)

            for (let x of relatedsub) {

                var colcount = $('.page0' + PageId).attr('data-collapcount')

                var Subpage = AddSubString(colcount, x.Name, 0, SubPageId);

                $('.page0' + PageId).children('.accordionOption').append(Subpage);

                $('.page0' + PageId).children('.spaceAccord').children('a').children('div').children('img').show();

                var obj = SubPageObjectCreate();

                obj.Name = x.Name

                obj.Content = x.Content

                obj.NewParentId = PageId

                newSubpagearr.push(obj)

                SubPageId++;
            }

        }

    }

    PageId++;

    OrderIndex++;

}

function FindSubPageArray(pageid, newpageid) {


    var newsubarray = []

    for (let x of updSubpagearr) {

        if (x.ParentId == pageid && x.NewParentId == newpageid) {

            newsubarray.push(x)

        }

    }

    for (let x of newSubpagearr) {

        if (x.ParentId == pageid && x.NewParentId == newpageid) {

            newsubarray.push(x)

        }

    }

    for (let x of tempSubpagearr) {

        if (x.ParentId == pageid && x.NewParentId == newpageid) {

            const pageindex = newsubarray.findIndex(obj => {

                return obj.SubPageId == x.SubPageId && obj.NewSubPageId == x.NewSubPageId;

            })

            if (pageindex < 0) {

                newsubarray.push(x)
            }



        }


    }


    console.log(newsubarray);

    return newsubarray
}

$(document).on('click', '.subpageclone', function () {

    var id = $(this).parents('.subpageview').attr('data-id');

    var newid = $(this).parents('.subpageview').attr('data-newid');

    SubPageClone(id, newid)


})

function SubPageClone(id, newid) {


    const updpageindex = updSubpagearr.findIndex(obj => {

        return obj.SubPageId == id && obj.NewSubPageId == newid;

    })

    const cpageindex = newSubpagearr.findIndex(obj => {

        return obj.SubPageId == id && obj.NewSubPageId == newid;

    })


    if (updpageindex >= 0) {

        var colcount = $(this).parents('.sort-index').attr('data-collapcount')

        var AddPage = AddSubString(colcount, updSubpagearr[updpageindex].Name, 0, SubPageId);

        $(AddPage).insertAfter($(".subpag" + id + newid));

        var pageobj = SubPageObjectCreate();

        pageobj.OrderIndex = OrderIndex

        pageobj.Name = updSubpagearr[updpageindex].Name

        pageobj.Content = updSubpagearr[updpageindex].Content

        pageobj.ParentId = updSubpagearr[updpageindex].ParentId

        pageobj.NewParentId = updSubpagearr[updpageindex].NewParentId

        pageobj.PageGroupId = parseInt($(this).parents('.grpss').attr('data-id'))

        pageobj.NewGroupId = parseInt($(this).parents('.grpss').attr('data-newid'))

        newSubpagearr.push(pageobj)

        SubPageId++;

        OrderIndex++;

    } else {

        if (cpageindex >= 0) {

            var colcount = $(this).parents('.sort-index').attr('data-collapcount')

            var AddPage = AddSubString(colcount, newSubpagearr[cpageindex].Name, 0, SubPageId);

            $(AddPage).insertAfter($(".subpag" + id + newid));

            var pageobj = SubPageObjectCreate();

            pageobj.OrderIndex = OrderIndex

            pageobj.Name = newSubpagearr[cpageindex].Name

            pageobj.Content = newSubpagearr[cpageindex].Content

            pageobj.ParentId = newSubpagearr[cpageindex].ParentId

            pageobj.NewParentId = newSubpagearr[cpageindex].NewParentId

            pageobj.PageGroupId = parseInt($(this).parents('.grpss').attr('data-id'))

            pageobj.NewGroupId = parseInt($(this).parents('.grpss').attr('data-newid'))

            newSubpagearr.push(pageobj)

            SubPageId++;

            OrderIndex++;
        } else {

            const temppageindex = tempSubpagearr.findIndex(obj => {

                return obj.SubPageId == id && obj.NewSubPageId == newid;

            })

            var colcount = $(this).parents('.sort-index').attr('data-collapcount')

            var AddPage = AddSubString(colcount, tempSubpagearr[temppageindex].Name, 0, SubPageId);

            $(AddPage).insertAfter($(".subpag" + id + newid));

            var pageobj = SubPageObjectCreate();

            pageobj.OrderIndex = OrderIndex

            pageobj.Name = tempSubpagearr[temppageindex].Name

            pageobj.Content = tempSubpagearr[temppageindex].Content

            pageobj.ParentId = tempSubpagearr[temppageindex].ParentId

            pageobj.NewParentId = tempSubpagearr[temppageindex].NewParentId

            pageobj.PageGroupId = parseInt($(this).parents('.grpss').attr('data-id'))

            pageobj.NewGroupId = parseInt($(this).parents('.grpss').attr('data-newid'))

            newSubpagearr.push(pageobj)

            SubPageId++;

            OrderIndex++;

        }


    }

}

$(document).on('keyup', '.groupnamee', function () {

    $(this).css('height', 'auto').css('height', this.scrollHeight + (this.offsetHeight - this.clientHeight));

})

/** Content Validation  */
function ContentValidation() {

    for (let x of newpagesarr) {

        console.log("x.Cpmtemt", x.Content)

        if (x.Content == '') {

            notify_content = `<div class="toast-msg dang-red" ><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>` + languagedata.Toast.pagecontent + `</span></div>`;
            $(notify_content).insertBefore(".header-rht");
            setTimeout(function () {
                $('.toast-msg').fadeOut('slow', function () {
                    $(this).remove();
                });
            }, 5000); // 5000 milliseconds = 5 seconds

            console.log(x);

            $('.page' + x.PageId.toString() + x.NewPageId.toString()).children('.spaceAccord').children('.spaceAccord-btn').trigger('click');

            return true
        }

        // if (x.ReadTime ==''){

        //     notify_content = `<div class="toast-msg dang-red" ><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>Please Enter the Reading Time of Content</span></div>`;
        //     $(notify_content).insertBefore(".header-rht");
        //     setTimeout(function () {
        //         $('.toast-msg').fadeOut('slow', function () {
        //             $(this).remove();
        //         });
        //     }, 5000); // 5000 milliseconds = 5 seconds



        //     $('.page' + x.PageId.toString() + x.NewPageId.toString()).children('.spaceAccord').children('.spaceAccord-btn').trigger('click');

        //     return true
        // }
     


    }

    for (let x of updPagesarr) {

        console.log("x.Cpmtemt", x.Content)

        if (x.Content == '') {

            notify_content = `<div class="toast-msg dang-red" ><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>` + languagedata.Toast.pagecontent + `</span></div>`;
            $(notify_content).insertBefore(".header-rht");
            setTimeout(function () {
                $('.toast-msg').fadeOut('slow', function () {
                    $(this).remove();
                });
            }, 5000); // 5000 milliseconds = 5 seconds

            $('.page' + x.PageId.toString() + x.NewPageId.toString()).children('.spaceAccord').children('.spaceAccord-btn').trigger('click');

            return true
        }
        // if (x.ReadTime ==''){

        //     notify_content = `<div class="toast-msg dang-red" ><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>Please Enter the Reading Time of Content</span></div>`;
        //     $(notify_content).insertBefore(".header-rht");
        //     setTimeout(function () {
        //         $('.toast-msg').fadeOut('slow', function () {
        //             $(this).remove();
        //         });
        //     }, 5000); // 5000 milliseconds = 5 seconds



        //     $('.page' + x.PageId.toString() + x.NewPageId.toString()).children('.spaceAccord').children('.spaceAccord-btn').trigger('click');

        //     return true
        // }
       


    }


    for (let x of newSubpagearr) {

        console.log("x.Cpmtemt", x.Content)

        if (x.Content == '') {

            notify_content = `<div class="toast-msg dang-red" ><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>` + languagedata.Toast.pagecontent + `</span></div>`;
            $(notify_content).insertBefore(".header-rht");
            setTimeout(function () {
                $('.toast-msg').fadeOut('slow', function () {
                    $(this).remove();
                });
            }, 5000); // 5000 milliseconds = 5 seconds

            $('.subpag' + x.SubPageId.toString() + x.NewSubPageId.toString()).trigger('click');

            return true
        }
        // if (x.ReadTime ==''){

        //     notify_content = `<div class="toast-msg dang-red" ><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>Please Enter the Reading Time of Content</span></div>`;
        //     $(notify_content).insertBefore(".header-rht");
        //     setTimeout(function () {
        //         $('.toast-msg').fadeOut('slow', function () {
        //             $(this).remove();
        //         });
        //     }, 5000); // 5000 milliseconds = 5 seconds



        //     $('.page' + x.PageId.toString() + x.NewPageId.toString()).children('.spaceAccord').children('.spaceAccord-btn').trigger('click');

        //     return true
        // }
     


    }


    for (let x of updSubpagearr) {

        console.log("x.Cpmtemt", x.Content)

        if (x.Content == '') {

            notify_content = `<div class="toast-msg dang-red" ><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span> ` + languagedata.Toast.pagecontent + `</span></div>`;
            $(notify_content).insertBefore(".header-rht");
            setTimeout(function () {
                $('.toast-msg').fadeOut('slow', function () {
                    $(this).remove();
                });
            }, 5000); // 5000 milliseconds = 5 seconds

            $('.subpag' + x.SubPageId.toString() + x.NewSubPageId.toString()).trigger('click');

            return true
        }

        // if (x.ReadTime ==''){

        //     notify_content = `<div class="toast-msg dang-red" ><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>Please Enter the Reading Time of Content</span></div>`;
        //     $(notify_content).insertBefore(".header-rht");
        //     setTimeout(function () {
        //         $('.toast-msg').fadeOut('slow', function () {
        //             $(this).remove();
        //         });
        //     }, 5000); // 5000 milliseconds = 5 seconds



        //     $('.page' + x.PageId.toString() + x.NewPageId.toString()).children('.spaceAccord').children('.spaceAccord-btn').trigger('click');

        //     return true
        // }
       

    }

    return false
}
function decodeHTML(html) {
    var txt = document.createElement("textarea");
    txt.innerHTML = html;
    return txt.value;
}
// x-min function//

$(document).on('keyup', '#readtimetag', function () {


    Readtime = $(this).val();

    // if( Readtime > 1 ){


    //     $('.minclass').text('mins')
    // }else{

    //     $('.minclass').text('min')

    // }

    var pageid = $('.pageselected').parents('.sort-index').attr('data-id');

    var newpageid = $('.pageselected').parents('.sort-index').attr('data-newid');

    var subpageid = $('.subpageselected').attr('data-id');

    var newsubpageid = $('.subpageselected').attr('data-newid');

    var value = $(this).val();

    if ($('.pageselected').hasClass('updpageview')) {

        if (pageid != 0 || newpageid != 0) {

            const updatepageindex = updPagesarr.findIndex(obj => {

                return obj.PageId == pageid && obj.NewPageId == newpageid;

            })
           
            if (updatepageindex >= 0) {

                

                if (updPagesarr[updatepageindex].ReadTime != Readtime) {

                    $('.page' + pageid + newpageid).children('.accordionOption').children('a').addClass('pageupdatedd');

                }
                console.log(updPagesarr,"checkupd")
                updPagesarr[updatepageindex].ReadTime = parseInt( Readtime)

            } else {


                const pageindex = tempPagesarr.findIndex(obj => {

                    return obj.PageId == pageid && obj.NewPageId == newpageid;

                })


                if (tempPagesarr[pageindex].ReadTime != Readtime) {

                    updPagesarr.push(tempPagesarr[pageindex]);

                    console.log(updPagesarr, "arry")

                    const updatepageindex1 = updPagesarr.findIndex(obj => {

                        return obj.PageId == pageid && obj.NewPageId == newpageid;

                    })

                    console.log(updatepageindex1, "")
                    if (updatepageindex1 >= 0) {

                        // if (updPagesarr[updatepageindex].Content != data || updPagesarr[updatepageindex].Name != $('#pagetitle').val()) {

                        $('.page' + pageid + newpageid).children('.accordionOption').children('a').addClass('pageupdatedd');

                        // }

                        updPagesarr[updatepageindex1].ReadTime = parseInt( Readtime)

                    }

                }


            }

        }

    } else {

        if (pageid != 0 || newpageid != 0) {

            const pageindex = newpagesarr.findIndex(obj => {

                return obj.PageId == pageid && obj.NewPageId == newpageid;

            })

            if (pageindex >= 0) {

                newpagesarr[pageindex].ReadTime = parseInt( Readtime)

            }

        }

        if (subpageid != 0 || newsubpageid != 0) {

            const pageindex = newSubpagearr.findIndex(obj => {

                return obj.SubPageId == subpageid && obj.NewSubPageId == newsubpageid;

            })

            if (pageindex >= 0) {

                newSubpagearr[pageindex].ReadTime = parseInt( Readtime)

            }

        }

    }

    if (subpageid != 0 || newsubpageid != 0) {

        const updatesubpageindex = updSubpagearr.findIndex(obj => {

            return obj.SubPageId == subpageid && obj.NewSubPageId == newsubpageid;

        })


        if (updatesubpageindex >= 0) {

            if (updSubpagearr[updatesubpageindex].ReadTime != Readtime ) {

                $('.subpag' + subpageid + newsubpageid).addClass('pageupdatedd');

            }
            updSubpagearr[updatesubpageindex].ReadTime = parseInt( Readtime)

        } else {

            const pageindex = tempSubpagearr.findIndex(obj => {

                return obj.SubPageId == subpageid && obj.NewSubPageId == newsubpageid;

            })

            if (pageindex >= 0) {

                // if (tempSubpagearr[pageindex].Content != data || tempSubpagearr[pageindex].Name != $('#pagetitle').val()) {

                updSubpagearr.push(tempSubpagearr[pageindex]);

                // }
            }
        }

    }

})



$(document).on('focus','#readtimetag',function(){

$('.space-readtime').addClass('focused')


})
$(document).on('blur','#readtimetag',function(){

    $('.space-readtime').removeClass('focused')
    
    
    })


 $(document).on('click','.space-readtime',function(){

    $(this).addClass('focused')
 })   


 function handleInputChange(event) {
    const target = event.currentTarget;
    const regex = /^[-+]?(?:[\d]+|(?:\.\d{1,3}))$/;
    
    // Replace the input with a valid number
    target.value = target.value.replace(/[^\d.-]/g, '').replace(/(\..*)\./g, '$1');
    
    // Ensure the new value matches the regular expression
    if (!regex.test(target.value)) {
      target.value = '';
    }
}