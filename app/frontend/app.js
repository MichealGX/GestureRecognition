document.addEventListener("DOMContentLoaded", function () {
  // 无需登录，直接显示主界面
  const mainContent = document.getElementById("mainContent");
  mainContent.style.display = "block";
});

// 控制侧边栏菜单展开/收起
function toggleMenu(menuId) {
const menu = document.getElementById(menuId);
menu.style.display = menu.style.display === "block" ? "none" : "block";
}
