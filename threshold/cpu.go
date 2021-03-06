package threshold

import "webup/syshealth"

type CPULoadTrigger struct {
}

func (trigger *CPULoadTrigger) GetKey() key {
	return "cpu.overload"
}

func (trigger *CPULoadTrigger) Check(metrics syshealth.Data) syshealth.ThresholdLevel {
	if rawLoad, ok := metrics["cpu.load_5"]; ok {
		if load, ok := rawLoad.(float64); ok {
			if load >= 0.8 {
				return syshealth.Critical
			}
			if load >= 0.6 {
				return syshealth.Warning
			}
		}
	}
	return syshealth.None
}
