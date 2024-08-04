layui.define([], function (exports) {

  const { admin, form, upload, $, element } = layui;

  function fetchUser() {
    admin.req({
      url: '/api/user',
      type: 'get',
      success: function (res) {
        form.val('profileForm', res.data);
        $('#upload-avatar-img').attr('src', res.data.avatar_url);
      }
    });
  }

  const uploadInst = upload.render({
    elem: '#upload-avatar-img',
    url: '/upload',
    before: function (obj) {
      // 预读本地文件示例，不支持ie8
      obj.preview(function (index, file, result) {
        $('#upload-avatar-img').attr('src', result);
      });
    },
    done: function (res) {
      if (res.code > 0) {
        return layer.msg('上传失败');
      }
      layer.msg('上传成功');

      form.val('profileForm', {
        avatar_url: res.data.url
      })

      $('#upload-avatar-text').html(''); // 置空上传失败的状态
    },
    error: function () {
      // 演示失败状态，并实现重传
      var demoText = $('#upload-avatar-text');
      demoText.html('<span style="color: #FF5722;">上传失败</span> <a class="layui-btn layui-btn-xs upload-reload">重试</a>');
      demoText.find('.upload-reload').on('click', function () {
        uploadInst.upload();
      });
    },
    // 进度条
    progress: function (n, elem, e) {
      element.progress('filter-upload-avatar', n + '%'); // 可配合 layui 进度条元素使用
      if (n == 100) {
        layer.msg('上传完毕', { icon: 1 });
      }
    }
  });

  //设置我的资料
  form.on('submit(updateProfile)', function (obj) {
    // layer.alert(layui.util.escape(JSON.stringify(obj.field)));
    const field = obj.field;
    field.gender = parseInt(field.gender);

    // 提交修改
    admin.req({
      url: '/api/user',
      method: 'PATCH',
      contentType: "application/json",
      data: JSON.stringify(field),
      success: function () {
        layer.msg('修改成功');
        fetchUser();
      }
    });

    return false;
  });

  fetchUser();

  exports('profile', {});
});
