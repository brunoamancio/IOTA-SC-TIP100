use wasmlib::{ScExports, ScViewContext};
use iota_sc_utils::{params, results, interfaces};

#[no_mangle]
pub fn on_load() {
    let exports = ScExports::new();
    exports.add_view(interfaces::NAME_FUNC_IMPLEMENTS, implements);
}

pub fn implements(ctx : &ScViewContext) {
    let hname_to_check = params::get_hname(interfaces::INTERFACE_TIP_100, ctx);
    let implements = hname_to_check == interfaces::HNAME_INTERFACE_TIP_100;
    results::set_bool(interfaces::INTERFACE_TIP_100, implements, ctx);
}