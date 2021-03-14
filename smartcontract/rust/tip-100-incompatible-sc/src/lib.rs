use iota_sc_utils::wasmlib::{ScExports, ScViewContext, ScBaseContext};

#[no_mangle]
pub fn on_load() {
    let exports = ScExports::new();
    exports.add_view("my_view_in_a_tip100_incompatible_contract", my_view_in_a_tip100_incompatible_contract);
}

pub fn my_view_in_a_tip100_incompatible_contract(ctx : &ScViewContext) {
    ctx.log("Hello world!");
}