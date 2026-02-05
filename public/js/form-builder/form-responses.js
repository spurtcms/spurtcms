$(document).ready(function () {

    const quill = new Quill('#editor', {
        theme: 'snow',
        modules: {
            toolbar: [
                ['bold', 'italic', 'underline'],
                [{ list: 'ordered' }, { list: 'bullet' }]
            ]
        }
    });


    $(".ql-snow").addClass("hidden")

    $("#replyButton").click(function () {
        $(".ql-snow").removeClass("hidden")
        $("#editor").removeClass("hidden")
        $("#editorSaveandCancel").removeClass("hidden")
        $("#replyButton").addClass("hidden")
    })

    $("#editorCancelBtn").click(function () {
        quill.setContents([]); // 🔥 clears editor

        $(".ql-snow").addClass("hidden");
        $("#editor").addClass("hidden");
        $("#editorSaveandCancel").addClass("hidden");
        $("#replyButton").removeClass("hidden");
    });

    $("#editorSubmitBtn").click(function () {

        const text = quill.getText().trim();

        if (!text) {

            notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  flex relative max-sm:max-w-[300px] items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-error.svg" alt = "toast error"></div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Warning</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >Please Enter the Content</p ></div ></div ></li></ul> `;
            $(notify_content).insertBefore(".editorerror");
            setTimeout(function () {
                $('.toast-msg').fadeOut('slow', function () {
                    $(this).remove();
                });
            }, 5000); // 5000 milliseconds = 5 seconds
            return;
        }

        let email = $("#editorEmail").val()

        let ticket = $("#ticketNumber").val()

        let username = $("#UserName").val()

        const htmlContent = quill.root.innerHTML;

        console.log("htmlContent::", htmlContent, "email::", email, "ticket::", ticket);

        $.ajax({
            url: "/admin/cta/replyforresponse",
            type: "POST",
            async: false,
            data: { csrf: $("input[name='csrf']").val(), "htmlContent": htmlContent, "email": email, "ticket": ticket, "username": username },
            datatype: "json",
            caches: false,
            success: function (data) {

                console.log("data:", data.status);

                if (data.status == true) {

                    setCookie("get-toast", "Reply Submitted Successfully")

                    window.location.href = "/admin/cta/form-responses"

                }
            }
        })


    })

    // Preview click
    $(document).on('click', '.responsepreview', function () {

        ctaid = $(this).attr('data-id')

        email = $(this).attr('data-email')

        ticket = $(this).attr('data-ticket')

        reply = $(this).attr('data-reply')

        username = $(this).attr('data-name')

        if (reply != "") {

            $("#replyButton").addClass("hidden")

        } else {
            $("#replyButton").removeClass("hidden")
        }

        $(".allResponses").addClass("hidden")

        $(".allReply").addClass("hidden")

        $("#response" + ctaid).removeClass("hidden")

        $("#reply" + ctaid).removeClass("hidden")

        $("#editorEmail").val(email)

        $("#ticketNumber").val(ticket)

        $("#UserName").val(username)


    })

    // Cancel the preview
    $("#cancel").click(function () {
        quill.setContents([]); // 🔥 clears editor

        $(".ql-snow").addClass("hidden");
        $("#editor").addClass("hidden");
        $("#editorSaveandCancel").addClass("hidden");
        $("#replyButton").removeClass("hidden");
        $(".allResponses").addClass("hidden")
    });

})
