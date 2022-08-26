// Code generated by 'ccgo -o pcre2_darwin_amd64.go -pkgname lib -trace-translation-units -export-externs X -export-defines D -export-fields F -export-structs S -export-typedefs T pcre.json .libs/libpcre2-8.a', DO NOT EDIT.

package lib

var CAPI = map[string]struct{}{
	"__darwin_check_fd_set_overflow":          {},
	"__isctype":                               {},
	"__istype":                                {},
	"__wcwidth":                               {},
	"_pcre2_OP_lengths_8":                     {},
	"_pcre2_auto_possessify_8":                {},
	"_pcre2_callout_end_delims_8":             {},
	"_pcre2_callout_start_delims_8":           {},
	"_pcre2_check_escape_8":                   {},
	"_pcre2_default_compile_context_8":        {},
	"_pcre2_default_convert_context_8":        {},
	"_pcre2_default_match_context_8":          {},
	"_pcre2_default_tables_8":                 {},
	"_pcre2_extuni_8":                         {},
	"_pcre2_find_bracket_8":                   {},
	"_pcre2_hspace_list_8":                    {},
	"_pcre2_is_newline_8":                     {},
	"_pcre2_jit_free_8":                       {},
	"_pcre2_jit_free_rodata_8":                {},
	"_pcre2_jit_get_size_8":                   {},
	"_pcre2_jit_get_target_8":                 {},
	"_pcre2_memctl_malloc_8":                  {},
	"_pcre2_ord2utf_8":                        {},
	"_pcre2_script_run_8":                     {},
	"_pcre2_strcmp_8":                         {},
	"_pcre2_strcmp_c8_8":                      {},
	"_pcre2_strcpy_c8_8":                      {},
	"_pcre2_strlen_8":                         {},
	"_pcre2_strncmp_8":                        {},
	"_pcre2_strncmp_c8_8":                     {},
	"_pcre2_study_8":                          {},
	"_pcre2_ucd_boolprop_sets_8":              {},
	"_pcre2_ucd_caseless_sets_8":              {},
	"_pcre2_ucd_digit_sets_8":                 {},
	"_pcre2_ucd_records_8":                    {},
	"_pcre2_ucd_script_sets_8":                {},
	"_pcre2_ucd_stage1_8":                     {},
	"_pcre2_ucd_stage2_8":                     {},
	"_pcre2_ucp_gbtable_8":                    {},
	"_pcre2_ucp_gentype_8":                    {},
	"_pcre2_unicode_version_8":                {},
	"_pcre2_utf8_table1":                      {},
	"_pcre2_utf8_table1_size":                 {},
	"_pcre2_utf8_table2":                      {},
	"_pcre2_utf8_table3":                      {},
	"_pcre2_utf8_table4":                      {},
	"_pcre2_utt_8":                            {},
	"_pcre2_utt_names_8":                      {},
	"_pcre2_utt_size_8":                       {},
	"_pcre2_valid_utf_8":                      {},
	"_pcre2_vspace_list_8":                    {},
	"_pcre2_was_newline_8":                    {},
	"_pcre2_xclass_8":                         {},
	"digittoint":                              {},
	"isalnum":                                 {},
	"isalpha":                                 {},
	"isascii":                                 {},
	"isblank":                                 {},
	"iscntrl":                                 {},
	"isdigit":                                 {},
	"isgraph":                                 {},
	"ishexnumber":                             {},
	"isideogram":                              {},
	"islower":                                 {},
	"isnumber":                                {},
	"isphonogram":                             {},
	"isprint":                                 {},
	"ispunct":                                 {},
	"isrune":                                  {},
	"isspace":                                 {},
	"isspecial":                               {},
	"isupper":                                 {},
	"isxdigit":                                {},
	"pcre2_callout_enumerate_8":               {},
	"pcre2_code_copy_8":                       {},
	"pcre2_code_copy_with_tables_8":           {},
	"pcre2_code_free_8":                       {},
	"pcre2_compile_8":                         {},
	"pcre2_compile_context_copy_8":            {},
	"pcre2_compile_context_create_8":          {},
	"pcre2_compile_context_free_8":            {},
	"pcre2_config_8":                          {},
	"pcre2_convert_context_copy_8":            {},
	"pcre2_convert_context_create_8":          {},
	"pcre2_convert_context_free_8":            {},
	"pcre2_converted_pattern_free_8":          {},
	"pcre2_dfa_match_8":                       {},
	"pcre2_general_context_copy_8":            {},
	"pcre2_general_context_create_8":          {},
	"pcre2_general_context_free_8":            {},
	"pcre2_get_error_message_8":               {},
	"pcre2_get_mark_8":                        {},
	"pcre2_get_match_data_size_8":             {},
	"pcre2_get_ovector_count_8":               {},
	"pcre2_get_ovector_pointer_8":             {},
	"pcre2_get_startchar_8":                   {},
	"pcre2_jit_compile_8":                     {},
	"pcre2_jit_free_unused_memory_8":          {},
	"pcre2_jit_match_8":                       {},
	"pcre2_jit_stack_assign_8":                {},
	"pcre2_jit_stack_create_8":                {},
	"pcre2_jit_stack_free_8":                  {},
	"pcre2_maketables_8":                      {},
	"pcre2_maketables_free_8":                 {},
	"pcre2_match_8":                           {},
	"pcre2_match_context_copy_8":              {},
	"pcre2_match_context_create_8":            {},
	"pcre2_match_context_free_8":              {},
	"pcre2_match_data_create_8":               {},
	"pcre2_match_data_create_from_pattern_8":  {},
	"pcre2_match_data_free_8":                 {},
	"pcre2_pattern_convert_8":                 {},
	"pcre2_pattern_info_8":                    {},
	"pcre2_serialize_decode_8":                {},
	"pcre2_serialize_encode_8":                {},
	"pcre2_serialize_free_8":                  {},
	"pcre2_serialize_get_number_of_codes_8":   {},
	"pcre2_set_bsr_8":                         {},
	"pcre2_set_callout_8":                     {},
	"pcre2_set_character_tables_8":            {},
	"pcre2_set_compile_extra_options_8":       {},
	"pcre2_set_compile_recursion_guard_8":     {},
	"pcre2_set_depth_limit_8":                 {},
	"pcre2_set_glob_escape_8":                 {},
	"pcre2_set_glob_separator_8":              {},
	"pcre2_set_heap_limit_8":                  {},
	"pcre2_set_match_limit_8":                 {},
	"pcre2_set_max_pattern_length_8":          {},
	"pcre2_set_newline_8":                     {},
	"pcre2_set_offset_limit_8":                {},
	"pcre2_set_parens_nest_limit_8":           {},
	"pcre2_set_recursion_limit_8":             {},
	"pcre2_set_recursion_memory_management_8": {},
	"pcre2_set_substitute_callout_8":          {},
	"pcre2_substitute_8":                      {},
	"pcre2_substring_copy_byname_8":           {},
	"pcre2_substring_copy_bynumber_8":         {},
	"pcre2_substring_free_8":                  {},
	"pcre2_substring_get_byname_8":            {},
	"pcre2_substring_get_bynumber_8":          {},
	"pcre2_substring_length_byname_8":         {},
	"pcre2_substring_length_bynumber_8":       {},
	"pcre2_substring_list_free_8":             {},
	"pcre2_substring_list_get_8":              {},
	"pcre2_substring_nametable_scan_8":        {},
	"pcre2_substring_number_from_name_8":      {},
	"toascii":                                 {},
	"tolower":                                 {},
	"toupper":                                 {},
}