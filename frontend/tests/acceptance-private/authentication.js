import { Selector } from "testcafe";
import testcafeconfig from "../acceptance/testcafeconfig";
import Page from "../page-model/page";
import Account from "../page-model/account";
import Settings from "../page-model/settings";
import Menu from "../page-model/menu";

fixture`Test authentication`.page`${testcafeconfig.url}`;

const page = new Page();
const account = new Account();
const menu = new Menu();
const settings = new Settings();

test.meta("testID", "authentication-001").meta({ type: "smoke" })("Login and Logout", async (t) => {
  await t.navigateTo("/browse");

  await t
    .expect(page.nameInput.visible)
    .ok()
    .expect(Selector(".input-search input").visible)
    .notOk();

  await t.typeText(page.nameInput, "admin", { replace: true });

  await t.expect(page.loginAction.hasAttribute("disabled", "disabled")).ok();

  await t.typeText(page.passwordInput, "photoprism", { replace: true });

  await t.expect(page.passwordInput.hasAttribute("type", "password")).ok();

  await t.click(page.togglePasswordMode);

  await t.expect(page.passwordInput.hasAttribute("type", "text")).ok();

  await t.click(page.togglePasswordMode);

  await t.expect(page.passwordInput.hasAttribute("type", "password")).ok();

  await t.click(page.loginAction);

  await t.expect(Selector(".input-search input", { timeout: 7000 }).visible).ok();

  await page.logout();

  await t
    .expect(page.nameInput.visible)
    .ok()
    .expect(Selector(".input-search input").visible)
    .notOk();

  await t.navigateTo("/settings");
  await t
    .expect(page.nameInput.visible)
    .ok()
    .expect(Selector(".input-search input").visible)
    .notOk();
});

test.meta("testID", "authentication-002").meta({ type: "smoke" })(
  "Login with wrong credentials",
  async (t) => {
    await page.login("wrong", "photoprism");
    await t.navigateTo("/favorites");

    await t
      .expect(page.nameInput.visible)
      .ok()
      .expect(Selector(".input-search input").visible)
      .notOk();

    await page.login("admin", "abcdefg");
    await t.navigateTo("/archive");

    await t
      .expect(page.nameInput.visible)
      .ok()
      .expect(Selector(".input-search input").visible)
      .notOk();
  }
);

test.meta("testID", "authentication-003").meta({ type: "smoke" })("Change password", async (t) => {
  await t.navigateTo("/browse");
  await page.login("admin", "photoprism");
  await t.expect(Selector(".input-search input", { timeout: 15000 }).visible).ok();
  await menu.openPage("settings");

  await t
    .click(settings.accountTab)
    .typeText(account.currentPassword, "wrong", { replace: true })
    .typeText(account.newPassword, "photoprism", { replace: true });

  await t.expect(account.confirm.hasAttribute("disabled", "disabled")).ok();

  await t.typeText(account.retypePassword, "photoprism", { replace: true });

  await t.expect(account.confirm.hasAttribute("disabled", "disabled")).notOk();

  await t
    .click(account.confirm)
    .typeText(account.currentPassword, "photoprism", { replace: true })
    .typeText(account.newPassword, "1234567", { replace: true })
    .typeText(account.retypePassword, "1234567", { replace: true });

  await t.expect(account.confirm.hasAttribute("disabled", "disabled")).ok();

  await t
    .typeText(account.currentPassword, "photoprism", { replace: true })
    .typeText(account.newPassword, "photoprism123", { replace: true });

  await t.expect(account.confirm.hasAttribute("disabled", "disabled")).ok();

  await t.typeText(account.retypePassword, "photoprism123", { replace: true });

  await t.expect(account.confirm.hasAttribute("disabled", "disabled")).notOk();

  await t.click(account.confirm);
  await page.logout();
  if (t.browser.platform === "mobile") {
    await t.wait(7000);
  }
  await page.login("admin", "photoprism");
  await t.navigateTo("/archive");

  await t
    .expect(page.nameInput.visible)
    .ok()
    .expect(Selector(".input-search input").visible)
    .notOk();

  await page.login("admin", "photoprism123");
  await t.expect(Selector(".input-search input").visible).ok();
  await menu.openPage("settings");

  await t
    .click(settings.accountTab)
    .typeText(account.currentPassword, "photoprism123", { replace: true })
    .typeText(account.newPassword, "photoprism", { replace: true })
    .typeText(account.retypePassword, "photoprism", { replace: true })
    .click(account.confirm);
  await page.logout();
  await page.login("admin", "photoprism");

  await t.expect(Selector(".input-search input").visible).ok();
  await page.logout();
});
