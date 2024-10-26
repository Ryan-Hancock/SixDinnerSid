'use client';

import React, { useState } from 'react';
import { Circle } from 'lucide-react';

const FeedingSchedule = () => {
    const [meals, setMeals] = useState([]);
    const [selectedDate, setSelectedDate] = useState(new Date());

    const mealTimes = ['Morning', 'Noon', 'Evening', 'Night'];
    const days = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'];

    const addMeal = async (day, mealTime, foodType) => {
        // Implementation for adding a meal
    };

    return (
        <div className="p-4">
            <div className="rounded-lg border bg-card text-card-foreground shadow-sm">
                <div className="flex flex-col space-y-1.5 p-6">
                    <h3 className="text-2xl font-semibold leading-none tracking-tight">
                        Cat Feeding Schedule
                    </h3>
                </div>
                <div className="p-6">
                    <div className="grid grid-cols-[100px_repeat(4,1fr)] gap-2">
                        {/* Header row with meal times */}
                        <div className="font-bold">Day</div>
                        {mealTimes.map(time => (
                            <div key={time} className="font-bold text-center">{time}</div>
                        ))}

                        {/* Schedule grid */}
                        {days.map(day => (
                            <React.Fragment key={day}>
                                <div className="font-medium">{day}</div>
                                {mealTimes.map(time => (
                                    <div key={`${day}-${time}`} className="border p-2 text-center rounded-md">
                                        <div className="flex gap-2 justify-center">
                                            <button
                                                className="flex items-center gap-1 hover:bg-slate-100 p-1 rounded"
                                                onClick={() => addMeal(day, time, 'biscuits')}
                                            >
                                                <Circle className="w-4 h-4" /> B
                                            </button>
                                            <button
                                                className="flex items-center gap-1 hover:bg-slate-100 p-1 rounded"
                                                onClick={() => addMeal(day, time, 'meat')}
                                            >
                                                <Circle className="w-4 h-4" /> M
                                            </button>
                                        </div>
                                        <input
                                            type="time"
                                            className="mt-1 w-full text-sm p-1 rounded border"
                                            onChange={(e) => {
                                                // Handle time change
                                            }}
                                        />
                                    </div>
                                ))}
                            </React.Fragment>
                        ))}
                    </div>
                </div>
            </div>
        </div>
    );
};

export default FeedingSchedule;