layui.extend({
  setter: "../config",
}).define(["setter"], function (exports) {
  const { form, layer, $, setter } = layui;
  // 提交事件
  form.on("submit(user-login)", function (data) {
    // 请求登入接口
    $.ajax({
      url: "/login",
      data: data.field,
      method: "POST",
      async: false,
      success: function (res) {
        console.log(res);
        layer.msg("登入成功", {
          icon: 1,
          time: 1000,
        }, function () {
          location.href = "/main.html";
        });
      },
    });
    return false; // 阻止默认 form 跳转
  });
  exports("login", {});
});
