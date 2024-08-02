layui.define([], function (exports) {

  const { $, form, layer, util } = layui;

  // 自定义验证规则
  form.verify({
    // 确认密码
    confirmPassword: function (value, item) {
      var passwordValue = $('#reg-password').val();
      if (value !== passwordValue) {
        return '两次密码输入不一致';
      }
    }
  });

  // 提交事件
  form.on('submit(demo-reg)', function (data) {
    var field = data.field; // 获取表单字段值

    // 是否勾选同意
    if (!field.agreement) {
      layer.msg('您必须勾选同意用户协议才能注册');
      return false;
    }

    // 显示填写结果，仅作演示用
    layer.alert(JSON.stringify(field), {
      title: '当前填写的字段值'
    });

    // 此处可执行 Ajax 等操作
    // …

    return false; // 阻止默认 form 跳转
  });

  // 普通事件
  util.on('lay-on', {
    // 获取验证码
    'reg-get-vercode': function (othis) {
      var isvalid = form.validate('#reg-cellphone'); // 主动触发验证，v2.7.0 新增
      // 验证通过
      if (isvalid) {
        layer.msg('手机号规则验证通过');
        // 此处可继续书写「发送验证码」等后续逻辑
        // …
      }
    }
  });
  exports('register', {});
});
