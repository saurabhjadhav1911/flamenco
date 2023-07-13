# AvailableJobSettingEvalOnSubmit

Enables the 'eval on submit' toggle button behavior for this setting. A toggle button will be shown in Blender's submission interface. When toggled on, the `eval` expression will determine the setting's value. Manually editing the setting is then no longer possible, and instead of an input field, the 'placeholder' string is shown. An example use is the to-be-rendered frame range, which by default automatically follows the scene range, but can be overridden manually when desired. 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**show_button** | **bool** | Enable or disable the &#39;eval on submit&#39; toggle button. | 
**placeholder** | **str** | Placeholder text to show when the manual input field is hidden (because eval-on-submit has been toggled on by the user).  | 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


